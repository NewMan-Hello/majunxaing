package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()    //获取路由引擎
	r.GET("/get", getMsg) //get方法
	r.Run(":9090")        //如果不指定IP地址、端口号，则默认为本机IP地址、8080端口
}

func getMsg(c *gin.Context) {
	t, _ := template.ParseFiles("E:\\goland\\WorkSpace\\ginProject\\src\\02.tpl")
	t.Execute(c.Writer, nil)

	name := c.Query("username") //获取 URL 中的数据
	//返回字符串类型数据
	//c.String(http.StatusOK, "欢迎您: %s", name)
	//返回 json 数据
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, //状态
		"msg":  "返回信息",        //描述信息
		"data": "欢迎您：" + name, //数据
	})

}
