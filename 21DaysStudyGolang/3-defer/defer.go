package main

import "fmt"

func a()  {
	fmt.Println("func A called...")
}

func b()  {
	fmt.Println("func B called...")
}

func c()  {
	fmt.Println("func C called...")
}



func main()  {
	defer a()
	defer b()
	defer c()
}
