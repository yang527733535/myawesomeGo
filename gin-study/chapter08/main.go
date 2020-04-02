package main

import "github.com/gin-gonic/gin"

//路由重定向
func main()  {

	r:=gin.Default()
	r.GET("/test", func(context *gin.Context) {

		context.Request.URL.Path = "/test2"
		r.HandleContext(context)
	})
	r.GET("/test2", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"213",
		})
	})
	r.Run(":8081")

}