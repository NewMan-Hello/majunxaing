package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main()  {
	http.HandleFunc("/",index)
	http.ListenAndServe("localhost:8888",nil)
}

func index(rw http.ResponseWriter,req *http.Request) {
	message := ""
	if req.Method == "POST" {
		req.ParseForm()
		if checkUsername(req.Form["username"][0]) {
			fmt.Fprintf(rw,"用户名：",req.Form["username"][0],"密码：",req.Form["password"][0])
			return
		} else {
			message = "用户名长度不符合要求"
		}
	}
	t, _ := template.ParseFiles("E:\\goland\\WorkSpace\\GoWorkSpace\\src\\index.tpl")
	t.Execute(rw,message)
}

func checkUsername(username string) bool {
	if len(username) >=6 && len(username)<=16 {
		return true
	}
	return false
}