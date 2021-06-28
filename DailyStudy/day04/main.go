package main

import "fmt"

//func main() {
//	list := new([]int)
//	l = append(list, 1)
//	fmt.Println(l)
//}

var (
	size    = 1024
	maxSize = size * 2
)

func main() {
	list1 := []int{1, 2, 3}
	list2 := []int{4, 5}
	//list1 = append(list1, list2)
	list1 = append(list1, list2...)
	fmt.Println(list1)

	fmt.Println(size, maxSize)
}
