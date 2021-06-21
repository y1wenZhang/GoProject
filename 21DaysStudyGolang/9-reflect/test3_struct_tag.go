package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `info:"json" doc:"female"`
	Age int `info:"age"`
}

func FindTag(input interface{}) {
	t := reflect.TypeOf(input)

	for i:=0; i < t.NumField();i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc\n")
		fmt.Println(taginfo, tagdoc)
	}
}

func main() {	
	var user User

	FindTag(user)
}
