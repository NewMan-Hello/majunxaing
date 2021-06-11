package main

import "fmt"

func main() {
	var n [10]int
	var i, j int

	//为数组 n 初始化元素
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	//遍历数组元素
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var l, m, k int
	//声明数组的同时快速初始化数组
	balance1 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//输出数组元素
	for l = 0; l < 5; l++ {
		fmt.Printf("balance1[%d] = %f\n", l, balance1[l])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//输出数组元素
	for m = 0; m < len(balance2); m++ {
		fmt.Printf("balance2[%d] = %f\n", m, balance2[m])
	}

	//将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	a := [...]int{9: 10}
	var p *[10]int = &a
	fmt.Println(p)

	x, y := 1, 2
	b := [...]*int{&x, &y}
	fmt.Println(b)

	q:= new([10]int)
	fmt.Println(q)
}
