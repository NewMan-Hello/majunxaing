package main

import "fmt"

func main() {
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("OK")

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL2
			}
			fmt.Println(i)
		}
	}
LABEL2:
	fmt.Println("OK")

LABEL3:
	for i := 0;i<10;i++{
		for {
			continue LABEL3
			fmt.Println(i)
		}
	}
	fmt.Println("OK")
}
