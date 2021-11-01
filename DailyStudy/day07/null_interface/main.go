package main

import (
	"fmt"
)

// 没有实现任何方法的接口类型
// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量。
// 关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。

// var xxx interface{}


func newInterface(a interface{}) interface{}{
	fmt.Printf("类型是:%T", a)
	return a
}

func justifyType(x interface{}){
	switch v := x.(type) {
	case string:
		fmt.Println("字符类型")
	case int:
		fmt.Println("int类型")
	case bool:
		fmt.Println("布尔类型")
	default:
		fmt.Printf("未知类型:%v", v)

	}
}

func main() {
	var a interface{}

	s := "hello"
	a = s
	fmt.Printf("a %s 类型是%T\n", a, a)
	b := false
	a = b
	fmt.Printf("a %t 类型是%T\n", a, a)
	i := 18
	a = i
	fmt.Printf("a %d 类型是%T\n", a, a)
	// 空接口的应用，1.作为函数的参数
	d := newInterface(a)
	fmt.Println(d)

	// 作为map的值
	var m = make(map[string]interface{})
	m["name"] = "zhanghang"
	m["age"] = 27
	m["height"] = 177
	fmt.Printf("%v", m)

	// 类型断言
	justifyType(b)
	justifyType(s)
	justifyType(i)

}