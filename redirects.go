package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	// 通过 POST 方法进行 HTTP 重定向。
	r.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	r.POST("/foo", func(c *gin.Context) {
		c.String(http.StatusOK, "foo")
	})

	// 路由重定向，使用 HandleContext
	r.GET("/test01", func(c *gin.Context) {
		c.Request.URL.Path = "/test02"
		r.HandleContext(c)
	})
	r.GET("/test02", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	r.Run(":8888")
}
