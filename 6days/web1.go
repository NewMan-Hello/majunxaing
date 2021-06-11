package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:1020", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if len(req.Form["name"]) > 0 {
		fmt.Fprint(rw, "Hello World", req.Form["name"][0])
	} else {
		fmt.Fprint(rw, "Hello World")
	}
}
