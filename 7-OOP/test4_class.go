package main

import "fmt"

type AnimalInterface interface {  // 类的多态
	Eat()
	Walk()
	Run()
}

type Dog struct {
	Name string
}

type Mouse struct {
	Name string
}

func (this *Dog) Eat() {
	fmt.Println("Dog is eating")
}

func (this *Dog) Walk() {
	fmt.Println("Dog is walking")
}

func (this *Dog) Run() {
	fmt.Println("Dog is running")
}

func (this *Mouse) Eat() {
	fmt.Println("Mouse is eating")
}

func (this *Mouse) Walk() {
	fmt.Println("Mouse is walking")
}

func (this *Mouse) Run() {
	fmt.Println("Mouse is running")
}


func main() {
	var animal AnimalInterface
	animal = &Dog{"Tom"}
	animal.Walk()
	animal.Eat()
	animal.Run()

	animal = &Mouse{"Jim"}
	animal.Walk()
	animal.Eat()
	animal.Run()
}
