package main

import "fmt"

type myInt int

type Book struct {
	title string
	auth string
}

func changeBook1(book Book)  {
	// 值传递
	book.title = "Java"

	fmt.Printf("change1 book value = %v\n", book)

}

func changeBook2(book *Book)  {
	// 指针传递
	book.title = "python"

	fmt.Printf("change2 book value = %v\n", book)
}

func main() {
	var num myInt

	fmt.Printf("num type = %T\n", num)

	var book Book
	book.title = "golang"
	book.auth = "zhanghang"

	fmt.Printf("book value = %v\n", book)

	changeBook1(book)
	fmt.Printf("main book value = %v\n", book)

	changeBook2(&book)
	fmt.Printf("main book value = %v\n", book)

}
