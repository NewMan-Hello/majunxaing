package main

import "fmt"

type I struct {
	Name string
}

type J struct {
	Name string
}

// Print 虽然 Go 语言中没有方法重载的概念，
//但是方法是和结构绑定的，绑定的结构不同，方法其实就不同了
func (i I) Print() {
	fmt.Println("I")
}

func (j J) Print() {
	fmt.Println("J")
}

type TZ int

func (tz *TZ) Increase(num int){
	*tz += TZ(num)
}

func main() {
	a := I{}
	a.Print() //method.value
	(I).Print(a) //method.expression

	b := J{}
	b.Print()

	var c TZ
	c.Increase(100)
	fmt.Println(c)
}
