package main

import (
	"fmt"
)

func main() {
	c := make(chan bool) //直接 make ，是双向通道，可存可取
	go func() {
		fmt.Println("GO Go Go !!!")
		c <- true //存
		close(c) //对 channel 类型进行迭代需要进行关闭
	}()

	for v := range c {
		fmt.Println(v)
	}

	/*
	有无缓存的区别：
	无缓存是一个同步阻塞的，如果不是空的不被读掉，会阻塞一直到读掉它
	有缓存是一个异步的，管自己运行，你读不赌他都不管你
	*/

}
