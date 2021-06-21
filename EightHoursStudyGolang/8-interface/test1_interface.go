package main

import "fmt"

type UserInterFace interface {
	Eat()
	Run()
	Swimming()
}

type User struct {
}

func (this *User) Eat() {
	fmt.Println("user eating")
}

func (this *User) Run() {
	fmt.Println("user running")
}

func (this *User) Swimming() {
	fmt.Println("user swimming")
}

func main() {
	var user UserInterFace
	user = &User{}
	user.Eat()
	user.Run()
	user.Swimming()
}
