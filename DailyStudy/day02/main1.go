package main

import "fmt"

func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key, val := range slice {
		//fmt.Println(&val)
		value := val
		m[key] = &value
	}

	for k,v := range m {
		fmt.Printf("Key:%d----->val:%d\n", k, v)
	}
}
