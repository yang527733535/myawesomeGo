package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":200,
			"message": "pong",
			"data":"123",
		})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080

}