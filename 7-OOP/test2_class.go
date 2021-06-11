package main

import "fmt"

type Hero struct {
	Name  string  // 类的封装
	Ad    int
	Level int
}

/*
func (this Hero) GetName() string {
	// 此时this为Hero的一个副本，值传递
	fmt.Println("hero name: ", this.Name)
	return this.Name
}

func (this Hero) SetName(newName string)  {
	this.Name = newName

	fmt.Println("hero name: ", this.Name)
}

func (this Hero) Show()  {
	fmt.Printf("value: %v\n", this)
}
*/

// 方法名和变量名，首字母大写表示公有方法和属性，外部可访问使用；首字母小写表示私有方法和属性，仅内部可访问

func (this *Hero) GetName() string {
	// 此时this为Hero的指针，引用传递
	fmt.Println("hero name: ", this.Name)
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName

	fmt.Println("hero name: ", this.Name)
}

func (this *Hero) Show() {
	fmt.Printf("value: %v\n", this)
}

func main() {
	hero := Hero{Name: "iron man", Ad: 100, Level: 10}

	hero.Show()

	hero.SetName("spider man")

	hero.GetName()
}
