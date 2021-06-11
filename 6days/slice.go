package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(s1)

	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s2 := a[5:10]
	s3 := a[5:len(a)]
	s4 := a[5:]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	s5 := make([]int,3,10)
	fmt.Println(s5,len(s5),cap(s5))
}
