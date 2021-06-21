package main

import "fmt"


func swap(a *int, b *int)  {
	var temp = *a
	*a = *b
	*b = temp
}


func main()  {
	var a = 10
	var b = 20

	swap(&a, &b)

	fmt.Println("a: ", a, "b: ", b)

	var p *int // 一级指针
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int  // 二级指针
	pp = &p
	fmt.Println(pp)
	fmt.Println(&p)

}
