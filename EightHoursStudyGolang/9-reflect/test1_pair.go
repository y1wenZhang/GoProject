package main

import "fmt"

type Reader interface{
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {

}

func (this *Book) ReadBook()  {
	fmt.Println("read a book")
}

func (this *Book) WriteBook()  {
	fmt.Println("write a book")
}

func main() {

	book := &Book{}

	var r Reader
	r = book

	var w Writer
	fmt.Printf("r type of %T\n", r)
	fmt.Printf("r type of %T\n", w)
	value, ok := r.(Writer)
	fmt.Println(value, ok)



}
