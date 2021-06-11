package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main()  {
	http.HandleFunc("/",upload)
	http.ListenAndServe("localhost:8080",nil)
}
func upload(rw http.ResponseWriter,req *http.Request) {
	message := ""
	if req.Method == "POST" {
		req.ParseMultipartForm(32<<20)
		file,handler,err := req.FormFile("upload")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(handler)
		defer file.Close()
		f, err := os.OpenFile("./"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
		message = "上传成功"
	}
	t, _ := template.ParseFiles("E:\\goland\\WorkSpace\\GoWorkSpace\\src\\upload.tpl")
	t.Execute(rw,message)
}
