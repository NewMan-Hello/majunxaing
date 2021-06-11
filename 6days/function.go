package main

import "fmt"

func main(){
	A("1",2)
	a := A
	a("3",4)

	//匿名函数
	b := func() {
		fmt.Println("AA")
	}
	b()

	//闭包
	f := B(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func A(b string,a ...int)  {
	fmt.Println(b,a)
}

func B(x int) func(int) int {
	return func(y int) int {
		return x+y
	}
}
