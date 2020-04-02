package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//获取参数 这个是通过URL来获取参数
//获取querystring参数
//querystring指的是URL中?后面携带的参数，例如：/user/search?username=小王子&address=沙河。 获取请求的querystring参数的方法如下：
func main()  {
	 r :=gin.Default()
	 r.GET("/user/search", func(c *gin.Context) {
		 username:=c.DefaultQuery("username","xixixi")
		 address:=c.Query("address")
		 //输出JSON给调用方
		 c.JSON(http.StatusOK,gin.H{
		 	"message":"OK",
		 	"username":username,
		 	"address":address,
		 })
	 })

	r.Run(":8081")
}
