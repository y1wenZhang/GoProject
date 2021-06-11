package main

import "fmt"

type Animal struct {
	Name string
	Ad 	 int
}

type Cat struct {
	Animal  // 类的继承
	color string
}

func (this *Animal) GetName() string {
	fmt.Println("name is ", this.Name)
	return this.Name
}

func (this *Animal) SetName(NewName string)  {
	this.Name = NewName
}

func (this *Cat) GetColor() string {
	return this.color
}


func main() {
	animal := Animal{"cat", 100}
	animal.GetName()

	cat := Cat{Animal{"tom", 100}, "red"}
	cat.GetName()
}
