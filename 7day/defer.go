package main

import "fmt"

func main() {
	//fmt.Println("a")
	//defer fmt.Println("b")
	//defer fmt.Println("c")

	//for i := 0; i < 3; i++ {
	//	defer fmt.Println(i)
	//}

	//for i := 0; i < 3; i++ {
	//	defer func() {
	//		fmt.Println(i)
	//	}()
	//}

	D()
	E()
	F()
}

func D() {
	fmt.Println("Func D")
}

func E() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in E")
		}
	}()
	panic("Panic in E")
}

func F() {
	fmt.Println("Func F")
}
