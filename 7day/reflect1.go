package main

import (
	"fmt"
	"reflect"
)

type Admin struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	Admin
	title string
}

func main() {
	m := Manager{Admin: Admin{1, "Ok", 12}, title: "123"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)

	a := Admin{1, "OK", 12}
	Set(&a)
	fmt.Println(a)
}

// Set 修改值
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	//判断是不是指针，判断能不能修改
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Id")
	//判断找没找到
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.Int{
		f.SetInt(2)
	}
}
