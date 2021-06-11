package main

import "fmt"

func main() {
	// slice的四种声明方式

	// 声明slice是一个切片，并且初始化，默认值为1，2，3
	//slice1 := []int{1, 2, 3}

	// 声明slice是一个切片，没有开辟空间，长度为0
	//var slice1 []int
	//slice1 = make([]int, 3)  // 开辟3个空间

	// 声明slice是一个切片，同时给slice分配3个空间，初始默认值都为0
	//var slice1 []int = make([]int, 3)

	// 声明slice是一个切片，同时给slice分配空间，3个空间，初始默认值为0，通过:=推导出slice是一个切片
	slice1 := make([]int, 3)

	fmt.Printf("length:%d, slice:%v", len(slice1), slice1)

	if slice1 == nil {
		fmt.Println("slice is null")
	} else {
		fmt.Println("slice is not null")
	}

}
