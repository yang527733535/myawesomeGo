package main

import "github.com/gin-gonic/gin"

//获取参数 请求的参数通过URL路径传递，例如：/user/search/小王子/沙河

func main()  {
	r:=gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username:=c.Param("username")
		address:=c.Param("address")
		c.JSON(200,gin.H{
			"message":"ok",
			"username":username,
			"address":address,
		})
	})
	r.Run(":8081")
}

