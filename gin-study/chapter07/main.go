package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	//HTTP 重定向很容易。 内部、外部重定向均支持。
	r :=gin.Default()
	r.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})

	r.Run(":8081")
}