package main

import (
	"fmt"
	"reflect"
)

// 反射
// 反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

// 任何接口值都由是一个具体类型和具体类型的值两部分组成的

// TypeOf
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf(" type is:%s", v)
}

//TypeName type name | type kind
func TypeName(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("name:%s kind:%s\n", v.Name(), v.Kind())
}

// ValueOf 通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Float32:
		fmt.Printf("type is float32, value is :%f\n", float32(v.Float()))
	case reflect.Int8:
		fmt.Printf("type is int,value is %d", int8(v.Int()))
	case reflect.String:
		fmt.Printf("type is string, value is %s", string(v.String()))
	}
}

// 通过反射设置值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())
	if v.Elem().Kind() == reflect.Int{
		v.Elem().SetInt(100)
	}
}

func structReflect(x interface{}){
	v := reflect.TypeOf(x)
	if v.Kind() == reflect.Struct{
		for i := 0; i < v.NumField();i++{
			field := v.Field(i)
			fmt.Printf("name: %s, index:%d, type:%s, json tag:%s\n", field.Name, i, field.Type, field.Tag)
		}
		//v.FieldByName()
	}
}

func structMethodReflect(x interface{}){
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if t.Elem().Kind() == reflect.Struct{
		fmt.Println(v.NumMethod())
		for i:=0; i<t.NumMethod();i++{
			method := v.Method(i)
			// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
			var args = []reflect.Value{}
			method.Call(args)
		}
	}
}

// 结构体反射
type person struct {
	name string `json:"name" int:"name"`
	age int `json:"age" ini:"age"`
}

func (p *person)Walk(){
	fmt.Println("person walk....")
}
func (p *person)Run(){
	fmt.Println("person run....")
}

func main() {
	//x := int8(10)
	//reflectType(x)
	//a := float32(1.234)
	//reflectType(a)
	//TypeName(x)
	//
	//type person struct {
	//	name string
	//	age  int
	//}
	//type student struct {
	//	name string
	//}
	//TypeName(person{name: "zhanghang", age: 8})
	//TypeName(student{name: "lisi"})
	//
	//reflectValue(x)
	//reflectValue(a)
	//
	//var m = 18
	//reflectSetValue(&m)
	//fmt.Println(m)
	p := &person{
		name: "张航",
		age: 27,
	}
	structReflect(p)
	structMethodReflect(p)
}
