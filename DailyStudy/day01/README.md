```go
package main

import "fmt"

func main() {
	deferCall()
}


func deferCall() {
	defer func(){ defer fmt.Println("打印前....")}()
	defer func(){ defer fmt.Println("打印中....")}()
	defer func(){ defer fmt.Println("打印后....")}()
	panic("触发异常")
}

```

输出为：
```bash
打印后....
打印中....
打印前....
panic: 触发异常
```

### 解析
defer的执行顺序是**后进先出**。当出现panic语句时，会先按照defer后进先出的顺序之后，最后执行panic

