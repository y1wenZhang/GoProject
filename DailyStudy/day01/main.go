package main

import "fmt"

func main() {
	deferCall()
}


func deferCall() {
	defer func(){ defer fmt.Println("打印前....")}()
	defer func(){ defer fmt.Println("打印中....")}()
	defer func(){ defer fmt.Println("打印后....")}()
	panic("触发异常")
}
