package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个带缓冲的channel
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行，发送的元素：", i, "len(c)=", len(c), ", cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("主go程接收元素：", num)
	}

	defer fmt.Println("主go程结束")
}
