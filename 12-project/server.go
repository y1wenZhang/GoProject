package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock sync.RWMutex

	// 消息广播的channel
	Message chan string
}


func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
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

	user := NewUser(conn)
	// 用户上线，将用户添加到OnlineMap中
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()

	// 当前用户上线了，进行广播
	s.BroadCast(user, "上线啦！！！")

	// 接收用户消息，进行广播
	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if n == 0 {
				s.BroadCast(user, "下线了")
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}

			// 提取用户消息
			msg := string(buf[:n-1])
			s.BroadCast(user, msg)
		}
	}()

	// 当前handler阻塞
	select {

	}
}

func (s *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.Ip, s.Port))
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
