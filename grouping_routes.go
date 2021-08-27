package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 简单的路由组： v1
	v1 := r.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	// 简单的路由组： v2
	v2 := r.Group("/v2")
	{
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.GET("/read", readEndpoint)
	}

	r.Run(":8888")
}

func loginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "loginEndpoint")
}

func submitEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "submitEndpoint")
}

func readEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "readEndpoint")
}
