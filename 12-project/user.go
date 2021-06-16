package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	Conn   net.Conn
	server *Server
}

// NewUser 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		Conn:   conn,
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
	// 用户下线
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()

	u.server.BroadCast(u, "下线了")
}

func (u *User) SendMsg(msg string) {
	u.Conn.Write([]byte(msg))
}

func (u *User) DoMessage(msg string) {

	if msg == "who" {
		// 查询当前在线用户
		for _, user := range u.server.OnlineMap {
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

			u.SendMsg("用户名修改成功, 用户名修改为：" + newName + "\n")
		}
	} else if len(msg) > 3 && msg[:3] == "to|" {
		toName := strings.Split(msg, "|")[1]
		message := strings.Split(msg, "|")[2]
		_, ok := u.server.OnlineMap[toName]
		if !ok {
			u.SendMsg("用户不存在，请输入who,查询在线用户\n")
		} else {
			u.SendMsg("发送成功\n")
			toUser := u.server.OnlineMap[toName]
			toUser.SendMsg(fmt.Sprintf("%s对你说:%s\n", u.Name, message))
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
