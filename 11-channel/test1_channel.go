package main

import "fmt"

func main() {
	// 定义一个channel
	c := make(chan int)

	go func(a, b int) {
		s := a + b
		c <- s // 将s发送给c channel
	}(10, 20)

	num := <-c // 从channel中接收数据

	fmt.Println("num = ", num)
}
