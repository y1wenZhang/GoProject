package main

import "fmt"

func printArray(myArray [4]int)  {
	// 固定数组，值传递
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100  // 值传递，无法修改原数组
}

func main()  {
	var myArray1 [10]int  // 固定长度的数组

	for index, value := range myArray1 {
		fmt.Println("index = ", index, "value = ",  value)
	}

	myArray2 := [4]int{1,2,3,4}
	for index, value := range myArray2 {
		fmt.Println("index = ",  index, "value = ", value)
	}

	printArray(myArray2)

	fmt.Println("==========")
	for index, value := range myArray2 {
		fmt.Println("index = ",  index, "value = ", value)
	}

}
