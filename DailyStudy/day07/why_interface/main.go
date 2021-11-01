package main

import "fmt"

// 为什么要实现interface类型

type cat struct{}
type dog struct{}

// 一个类型可以实现多个接口
// 多个类型也可实现一个接口
type sayer interface {
	say()
}

type mover interface {
	move()
}

// 嵌套
/*type animals interface {
	sayer
	mover
}*/

type animals interface {
	say()
	move()
}

func (c *cat) say() {
	fmt.Println("汪汪汪...")
}
func (d *dog) say() {
	fmt.Println("喵喵喵...")
}

func(d *dog)move(){
	fmt.Println("跑.....")
}

func main() {
	var s sayer
	d := &dog{}
	s = d
	s.say()
	c := &cat{}
	s = c
	s.say()

	var a animals
	a = d
	a.say()
	a.move()
}
