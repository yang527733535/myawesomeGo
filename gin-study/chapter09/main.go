package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {

	r:=gin.Default()
	r.GET("/test", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"get",
		})
	})

	r.Any("/test2", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"Any",
		})
	})
	r.Run(":8081")
}
//r.NoRoute(func(c *gin.Context) {
//	c.HTML(http.StatusNotFound, "views/404.html", nil)
//})
//为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回views/404.html页面。