package main

import (
	"fmt"
)

//声明全局变量
var g int = 20

func main() {
	//声明局部变量
	var g int = 10
	fmt.Println(g)

	//声明局部变量
	var a int
	a = 10
	var b = 20
	c := 0
	fmt.Printf("main() 函数中 a = %d\n", a)
	c = sum(a, b)
	fmt.Printf("main() 函数中 c = %d\n", c)
}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}
