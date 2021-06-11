package main

import "fmt"

func main() {
	sum1 := 0
	for i := 0; i <= 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	sum2 := 1
	for ; sum2 <= 10; {
		sum2 += sum2
	}
	fmt.Println(sum2)

	for sum2 <= 10 {
		sum2 = +sum2
	}
	fmt.Println(sum2)

	strings := []string{"google", "baidu"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}
