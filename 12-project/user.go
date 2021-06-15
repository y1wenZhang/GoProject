package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C chan string
	Conn net.Conn
	server *Server
}

// NewUser 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		Conn: conn,
		server: server,
	}
	// 启动监听
	go user.ListenMessage()

	return user
}

func (u *User) Online() {
	// 用户上线，将用户添加到OnlineMap中
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()

	// 当前用户上线了，进行广播
	u.server.BroadCast(u, "上线啦！！！")
}

func (u *User) Offline() {
	u.server.BroadCast(u, "下线了")
}

func (u *User) SendMsg(msg string) {
	u.Conn.Write([]byte(msg))
}

func (u *User) DoMessage(msg string) {

	if msg == "who" {
		// 查询当前在线用户
		for _,user := range u.server.OnlineMap {
			msg := fmt.Sprintf("%s==%s在线\n", user.Addr, user.Name)
			u.SendMsg(msg)
		}
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 用户重命名
		newName := strings.Split(msg, "|")[1]
		_, ok := u.server.OnlineMap[newName]
		if ok {
			u.SendMsg("用户名已存在\n")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.server.OnlineMap[newName] = u
			u.server.mapLock.Unlock()

			u.Name = newName

			u.SendMsg("用户名修改成功\n")
		}
	} else {
		u.server.BroadCast(u, msg)
	}
}

// ListenMessage 监听当前User channel的方法，一旦有消息，就直接发送给对端客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.Conn.Write([]byte(msg + "\n"))
	}
}



