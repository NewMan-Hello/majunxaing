package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() //获取路由引擎
	r.GET("/GetOtherData", func(c *gin.Context) {
		url := "https://www.baidu.com/s?ie=utf-8&wd=%E7%A5%9E%E8%88%9F%E5%8D%81%E4%BA%8C%E5%8F%B7%E8%88%AA%E5%A4%A9%E5%91%98%E5%90%8D%E5%8D%95%E7%A1%AE%E5%AE%9A&tn=88093251_22_hao_pg&"
		response, err := http.Get(url)
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable) //应答 client
			return
		}
		body := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		//数据写入响应体
		c.DataFromReader(http.StatusOK,contentLength,contentType,body,nil)
	})
	r.Run(":9090")
}