package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() //获取路由引擎

	//一般重定向:重定向到外部网络
	r.GET("redirect1",func(c *gin.Context) {
		//重定向到百度，获取百度对应的数据
		//重定向状态码：StatusMovedPermanently
		url := "http://www.baidu.com" //重定向的 URL
		c.Redirect(http.StatusMovedPermanently,url)
	})

	//路由重定向:重定向到具体的路由
	r.GET("/redirect2", func(c *gin.Context) {
		c.Request.URL.Path = "/TestRedirect" //重定向的路由
		r.HandleContext(c)
	})

	//路由：localhost:9090/TestRedirect
	r.GET("/TestRedirect", func(c *gin.Context) {
		//返回 json 数据
		c.JSON(http.StatusOK,gin.H{
			"code": http.StatusOK, //状态
			"msg": "响应成功", //描述信息
		})
	})
	r.Run(":9090") //本机IP地址，9090端口
}
