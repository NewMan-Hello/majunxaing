package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()      //获取路由引擎
	r.GET("/post", postMsg) //post方法
	r.Run(":9090")          //本机IP地址，9090端口
}

func postMsg(c *gin.Context) {
	t, _ := template.ParseFiles("E:\\goland\\WorkSpace\\ginProject\\src\\03.tpl")
	t.Execute(c.Writer, nil)

	name := c.DefaultPostForm("username", "Gin") //获取 body 中的数据（ form 形式）
	fmt.Println(name)
	form, b := c.GetPostForm("username") //获取 body 中的数据（ form 形式）
	fmt.Println(form, b)
	c.JSON(http.StatusOK, "欢迎您："+name)
}
