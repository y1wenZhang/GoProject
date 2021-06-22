package main

import "fmt"

func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key, value := range slice {
		m[key] = &value
	}
	for k,v := range m {
		fmt.Printf("key:%d--->value:%d\n", k,*v)
	}
}
