package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// ListenMessage  监听Message广播消息channel的goroutine，一旦有消息发送给全部在线的User
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message

		// 将消息发送给全部在线用户
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// BroadCast 广播消息的方法
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]===##%s##===:%s", user.Addr, user.Name, msg)
	s.Message <- sendMsg
}

// Handler 业务处理
func (s *Server) Handler(conn net.Conn) {
	fmt.Println("连接服务器成功")

	user := NewUser(conn, s)
	user.Online()

	isAlive := make(chan bool)
	// 接收用户消息，进行广播
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}

			// 提取用户消息
			msg := string(buf[:n-1])
			user.DoMessage(msg)

			isAlive <- true
		}
	}()

	// 当前handler阻塞
	for {
		select {
		case <-isAlive:
			// 当前用户是活跃的，重置定时值
			// 不做任何事情，为了激活select，更新下面的定时器
		case <-time.After(10 * time.Second):  // 判断该条件是触发重置定时
			// 已经超时
			// 将当前User强制关闭
			user.SendMsg("你被踢了")

			// 销毁用户资源
			close(user.C)

			// 关闭连接
			conn.Close()

			// 退出当前Handler
			runtime.Goexit()

		}
	}
}

func (s *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net listern error: ", err)
	}
	// listen close
	defer listener.Close()

	// 启动监听Message的goroutine
	go s.ListenMessage()
	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error: ", err)
			continue
		}
		// do handler
		go s.Handler(conn)
	}

}
