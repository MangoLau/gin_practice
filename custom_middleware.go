package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 自定义中间件

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main_middleware() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// 打印： "12345"
		log.Println(example)
		c.String(http.StatusOK, example)
	})

	r.Run(":8888")
}
