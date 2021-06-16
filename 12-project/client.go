package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp string
	ServerPort int
	Name string
	Conn net.Conn
	Param int
}


func (client *Client) Menu() bool {
	var param int
	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 更新用户名")
	fmt.Println("0. 退出")

	fmt.Scanln(&param)

	if param >= 0 && param <=3 {
		client.Param = param
		return true
	} else {
		fmt.Println(">>>>>>>>请输入有效的数字<<<<<<<")
		return false
	}
}


func (client *Client) Run() {
	for client.Param != 0 {
		for client.Menu() != true {
		}
		switch client.Param {
		case 1:
			fmt.Println(">>>>>>>>公聊模式选择中<<<<<<")
		case 2:
			fmt.Println(">>>>>>>>私聊模式<<<<<<")
		case 3:
			fmt.Println(">>>>>>>>更改用户名<<<<<<")
		case 0:
			fmt.Println(">>>>>>>>退出<<<<<<<<<")
		}
	}
}


func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp: serverIp,
		ServerPort: serverPort,
		Param: 999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net Dial error:", err)
	}
	client.Conn = conn

	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>>CONNECT SERVER FAILED<<<<<<<<<<")
	}
	fmt.Println(">>>>>>>>>>CONNECT SERVER SUCCESS<<<<<<<<<<<")

	client.Run()
}
