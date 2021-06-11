package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := strconv.Itoa(a)
	fmt.Println(b)
	a, _ = strconv.Atoi(b)
	fmt.Println(a)
}