package main

import "fmt"

func main() {
	c := make(chan int, 3) // 未向其中传入数据，则为nil channel

	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}

		close(c) // close 可以关闭一个channel
		// channel 不像文件那样需要经常去关闭， 只要当你确实没有任何发送数据了，或者你想显示的结束range循环之类的，才会关闭channel
		// 关闭channel后无法再向其发送数据(会引发panic异常，导致接收立即返回零值)
		// 关闭channel可以继续从中获取数据
		// 对于nil channel，无论收发都会阻塞
	}()
	// fatal error: all goroutines are asleep - deadlock!   死锁异常
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	defer fmt.Println("main finished...")
}
