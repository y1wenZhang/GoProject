package main

import (
	"flag"
	"fmt"
	"io"
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


func (client *Client) PublicChat() {
	var msg string
	fmt.Println(">>>>>>>请输入聊天内容，exit退出：")
	fmt.Scanln(&msg)

	for msg != "exit" {
		if len(msg) != 0 {
			sendMsg := msg + "\n"
			_, err := client.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("client conn write error:", err)
				break
			}
		}

		msg = ""
		fmt.Println(">>>>>>>请输入聊天内容，exit退出：")
		fmt.Scanln(&msg)
	}
}


func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>>>请输入用户名：")
	fmt.Scanln(&client.Name)

	sendMsg := fmt.Sprintf("rename|%s", client.Name)

	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("client conn write error:", err)
		return false
	}
	return true
}


func (client *Client) DealResponse() {
	for {
		buf := make([]byte, 4096)
		n, err := client.Conn.Read(buf)
		if n == 0 {
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("Conn Read err: ", err)
			return
		}

		msg := string(buf[:n-1])
		fmt.Println(msg)
		//io.Copy(os.Stdout, client.Conn)
	}
}


func (client *Client) Run() {
	for client.Param != 0 {
		for client.Menu() != true {
		}
		switch client.Param {
		case 1:
			client.PublicChat()
		case 2:
			fmt.Println(">>>>>>>>私聊模式<<<<<<")
		case 3:
			client.UpdateName()
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

	go client.DealResponse()

	if client == nil {
		fmt.Println(">>>>>>>>>CONNECT SERVER FAILED<<<<<<<<<<")
	}
	fmt.Println(">>>>>>>>>>CONNECT SERVER SUCCESS<<<<<<<<<<<")

	client.Run()
}
