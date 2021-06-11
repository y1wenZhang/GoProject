package main

import "fmt"

func main() {

	// =========切片的容量，切片扩容==========
	// make(切片类型，切片长度，切片容量)
	var slice = make([]int, 3, 5)
	fmt.Printf("length: %d, cap: %d, value=%v\n", len(slice), cap(slice), slice)

	slice = append(slice, 1)
	fmt.Printf("length: %d, cap: %d, value=%v\n", len(slice), cap(slice), slice)


	// 默认切片容量=切片长度
	var slice1 = make([]int, 3)

	// 当切片长度大于当前切片容量时，会重新开辟一个新的切片容量
	slice1 = append(slice1, 1)
	fmt.Printf("length: %d, cap: %d, value=%v\n", len(slice1), cap(slice1), slice1)

	// 切片长度和切片容量不同，切片长度是指左指针到右指针的距离，切片容量是指左指针到底层数组末尾的距离

	// ===================切片的截取==============
	slice2 := []int{1,2,3}
	s := slice2[0:2]  // 元素索引，左闭右开区间
	s[0] = 100 // 截取的切片仍指向原切片的地址，修改其中的某个子元素，原切片也会更改
	fmt.Println(s)
	fmt.Println(slice2)

	// copy 函数，可以将切片的底层数组slice一起拷贝，深拷贝
	slice3 := []int{1,2,3}
	slice4 := make([]int, 3)  // 默认值[0,0,0]
	copy(slice4, slice3)  // 一次将slice3中的值拷贝到slice4中
	fmt.Println(slice4)
	slice4[0] = 100
	fmt.Println(slice4)  // copy函数，将slice3中的数字拷贝到slice4中，slice4指向的是slice，修改子元素指向的是slice4的地址
	fmt.Println(slice3)
}
