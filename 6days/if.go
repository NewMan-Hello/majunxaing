package main

import "fmt"

func main() {
	//定义局部变量
	var a int = 10
	b := 20
	//使用 if 语句判断布尔表达式
	if a < 20 {
		//如果条件为 true ，则执行以下语句
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	//判断布尔表达式
	if a < 10 {
		//如果条件为 true ，则执行以下语句
		fmt.Printf("a 小于 10\n")
	} else {
		//如果条件为 false ，则执行以下语句
		fmt.Printf("a 不小于 10\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	//判断条件
	if a < 20 {
		// if 条件语句为 true 执行
		if b < 30 {
			// if 条件语句为 true 执行
			fmt.Printf("a 的值为: %d , b 的值为: %d\n", a, b)
		}
	}
}
