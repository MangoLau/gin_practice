package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 设置和获取 Cookie

func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 300, "/", "localhost", false, true)
		}

		c.String(http.StatusOK, cookie)
	})

	r.Run(":8888")
}
