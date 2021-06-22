### 下面这段代码输出什么，说明原因。

```go
package main

import "fmt"

func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key, value := range slice {
		m[key] = &value
	}
	for k,v := range m {
		fmt.Printf("key:%d--->value:%d\n", k,*v)
	}
}
```
**输出结果**
```bash
key:1--->value:3
key:2--->value:3
key:3--->value:3
key:0--->value:3
```

- 解析 
  
for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.

- 正确的写法
```go
package main

import "fmt"

func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key, val := range slice {
		//fmt.Println(&val)
		value := val
		m[key] = &value
	}

	for k,v := range m {
		fmt.Printf("Key:%d----->val:%d\n", k, v)
	}
}
```