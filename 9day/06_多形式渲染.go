package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// json 格式输出
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"html": "<b>hello,Gin框架</b>",
		})
	})

	//原样输出 html （ html 渲染）
	r.GET("/html", func(c *gin.Context) {
		c.PureJSON(http.StatusOK,gin.H{
			"html": "<b>hello,Gin框架</b>",
		})
	})

	//输出 xml 形式（ XML 渲染）
	r.GET("/xml", func(c *gin.Context) {
		//定义一个结构
		type Message struct {
			Name string
			Msg string
			Age int
		}
		//初始化
		info := Message{}
		info.Name = "小马"
		info.Msg = "hello"
		info.Age = 24
		c.XML(http.StatusOK,info)
	})

	//返回 yaml 形式（ yaml 渲染）
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK,gin.H{
			"message": "Gin框架的多形式渲染",
			"status": 200,
		})
	})

	//开启服务
	r.Run(":9090")
}
