package main

import (

	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/ping：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/ping",mytest)


	//JSON渲染
	r.GET("/someJSON", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})

	r.GET("/some2JSON", func(c *gin.Context) {
		// 方法二：使用结构体

		var msg struct{
			Name  string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "yangtenghui"
		msg.Message = "xixi"
		msg.Age = 12
		c.JSON(http.StatusOK,msg)
	})


	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	r.Run(":8081") // listen and serve on 0.0.0.0:8080

}

func mytest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
