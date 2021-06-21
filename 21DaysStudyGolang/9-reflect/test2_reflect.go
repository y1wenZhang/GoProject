package main

import (
	"fmt"
	"reflect"
)

type AnUser struct {
	Id   int
	Name string
	Age  int
}

func (this AnUser) Call()  {
	fmt.Println("user is ", this.Name)
}


func DoField(input interface{}) {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i:=0; i<inputType.NumMethod();i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}

func main() {
	user := AnUser{10, "yiwenzhang", 18}
	DoField(user)
}

