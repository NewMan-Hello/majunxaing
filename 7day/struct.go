package main

import "fmt"

//type person struct {
//}

type person struct {
	Name string
	Age  int
}

type people struct {
	Name string
	Age  int
	//匿名结构嵌套在其他结构
	Contact struct {
		Phone, City string
	}
}

type man struct {
	//匿名字段
	string
	int
}

// G 值拷贝
func G(per person) {
	per.Age = 13
	fmt.Println("G", per)
}

// H 地址拷贝
func H(per *person) {
	per.Age = 13
	fmt.Println("H", per)
}

type human struct {
	Sex string
}
type teacher struct {
	human //结构名称作为字段名称，默认其中字段都给了 teacher 结构
	Name  string
	Age   int
}
type student struct {
	human
	Name string
	Age  int
}

func main() {
	//test := person{}
	//fmt.Println(test)

	a := person{
		Name: "Joe",
		Age:  19,
	}
	//a.Name = "Joe"
	//a.Age = 19
	fmt.Println(a)

	G(a)
	fmt.Println(a)

	H(&a)
	fmt.Println(a)

	//匿名结构
	b := struct {
		Name string
		Age  int
	}{
		Name: "Joe",
		Age:  19,
	}
	fmt.Println(b)

	//c := people{}
	//fmt.Println(c)
	c := people{Name: "Joe", Age: 19}
	c.Contact.Phone = "12345678911"
	c.Contact.City = "beijing"
	fmt.Println(c)

	d := man{"Joe", 19}
	fmt.Println(d)

	e := a
	fmt.Println(e)

	f := teacher{Name: "Joe", Age: 19, human: human{Sex: "男"}}
	f.Sex = "女"
	g := teacher{Name: "Joe", Age: 20, human: human{Sex: "女"}}
	fmt.Println(f, g)
}
