package main

import (
	"fmt"
	"time"
)

// NewTask 子goroutine
func NewTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {

	go NewTask()
	// 主goroutine终止，子goroutine也将会终止
	i := 0
	for {
		i++
		fmt.Printf("main Goroutine i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
