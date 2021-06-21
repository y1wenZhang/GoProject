package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/*
		// 匿名函数
			func(){ // 不带有形参的匿名函数

			}()  //在函数的末尾添加小括号()， 则表示执行该函数
	*/

	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// return  // 在子函数中使用return，只能退出当前子函数的执行
			runtime.Goexit() // 在子函数中使用Goexit，可以直接退出子函数所在的Goroutine
			fmt.Println(" func B")
		}()
		//return  // return 只能退出当前函数(如果是go协程，则退出协程)

		fmt.Println("func A")
	}()

	// 携带形参的匿名函数
	go func(a, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20) // 通过在末尾的的括号中传入参数

	for {
		time.Sleep(1 * time.Second)
	}
}
