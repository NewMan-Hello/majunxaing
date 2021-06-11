package main

import "fmt"

func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)

	m[1] = make(map[int]string)
	m[1][1] = "OK"
	a := m[1][1]
	fmt.Println(a)

	b, ok := m[2][1]
	if !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "GOOD"
	b, ok = m[2][1]
	fmt.Println(b, ok)
}
