package main

import (
	"fmt"
	"reflect"
)

type AUser struct {
	Id   int
	Name string
	Age  int
}

func (u AUser) Hello1(name string) {
	fmt.Println("Hello",name,",my name is",u.Name)
}

func main(){
	u := AUser{1,"Ok",12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello1")

	//建立参数，要求传入一个 slice
	args := []reflect.Value{reflect.ValueOf("Joe")}
	mv.Call(args)
}