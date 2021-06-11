package main

import "fmt"

func changeValue(myMap map[string]string)  {
	// 引用传递
	fmt.Println("before change: ", myMap)

	myMap["one"] = "shanghai"

	fmt.Println("func after change: ", myMap)

}

func main() {
	// ===第一种 map的声明方式
	// 声明map1是map类型，key为string，value为string
	//var map1 = make(map[string]string)
	//map1["One"] = "China"
	//fmt.Println(map1)

	var map1 map[string]string
	// make 为map开辟空间
	map1 = make(map[string]string, 10)

	fmt.Println(map1)
	// ===第二种 map的声明方式
	map2 := make(map[string]string)
	map2["Two"] = "Beijing"
	fmt.Println(map2)

	fmt.Println(map2["Two"])

	// ===第三种 map的声明方式
	map3 := map[string]string {
		"one": "beijing",
		"two": "shanghai",
		"three": "shenzhen",
	}
	fmt.Println(map3)

	// =======map的操作=========
	map4 := make(map[string]string)
	// 增
	map4["one"] = "beijing"
	map4["two"] = "shanghai"
	map4["three"] = "shenzhen"

	fmt.Println("add value", map4)

	// 删
	delete(map4, "one")
	fmt.Println("delete value", map4)

	// 改
	map4["two"] = "shenzhen"
	fmt.Println("update value", map4)

	// 查
	fmt.Println("get value: ", map4)

	// 在函数内部修改map内的值
	map5 := map[string]string {
		"one": "beijing",
		"two": "shanghai",
	}
	changeValue(map5)

	fmt.Println("out  after func change: ", map5)

}
