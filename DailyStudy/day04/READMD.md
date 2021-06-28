
#### 能否通过编译，不能的话原因是什么；如果能，输出什么。

```go
package main

import "fmt"

func main() {
	list := new([]int)
	l = append(list, 1)
	fmt.Println(l)
}
```
- 解析
  
不能通过编译，**`new([]int)` 之后的 list 是一个 `*[]int` 类型的指针，不能对指针执行 append 操作**。可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。

***

#### 能否通过编译，不能的话原因是什么；如果能，输出什么。

```go
func main() {
	list1 := []int{1, 2, 3}
	list2 := []int{4,5}
	list1 = append(list1, list2)
	fmt.Println(list1)
}
```
- 解析

不能通过编译，`append()` 的第二个参数不能直接使用 slice，需使用 `...` 操作符，将一个切片追加到另一个切片上：`append(s1,s2…)`。或者直接跟上元素，形如：`append(s1,1,2,3)`。
***

```go
var (
    size := 1024
    max_size = size*2
)

func main() {
    fmt.Println(size,max_size)
}
```
- 解析

不能通过编译。这道题的主要知识点是变量声明的简短模式，形如：x := 100。但这种声明方式有限制：
```text
1.必须使用显示初始化；
2.不能提供数据类型，编译器会自动推导；
3.只能在函数内部使用简短模式；
```
***