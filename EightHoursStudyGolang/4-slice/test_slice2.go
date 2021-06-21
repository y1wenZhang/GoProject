package main

import "fmt"


func printArray1(myArray []int)  {
	// 动态数组，引用传递
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100  // 引用传递，值修改，原数组值也修改
}

func main() {
	var myArray []int  // 动态数组
	fmt.Printf("myArray type = %T\n", myArray)

	myArray2 := []int{1,2,3,4}

	printArray1(myArray2)

	for _, value := range myArray2 {
		fmt.Println("value = ", value)
	}

}
