package main

import "fmt"

/*// 执行到 index 为 9 就结束
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //多核
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	<- c
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	if index == 9 {
		c <- true
	}
}*/

/*// 解决方案1
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //多核
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}*/

/*// 解决方案2 ，通过 sync 包
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //多核
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}

	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}*/

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	close(c2)

	<- o
}
