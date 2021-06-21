package main

import "fmt"

func funcDefer() int {
	fmt.Println("func defer called...")
	return 0
}

func funcReturn() int  {
	fmt.Println("func return called...")
	return 0
}


func returnAndDefer() int {
	defer funcDefer()
	return funcReturn()
}


func main() {
	returnAndDefer()  // return后的语句先执行，defer后的语句后执行
}
