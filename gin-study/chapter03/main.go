package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//获取参数 这个是通过Body来获取参数

func main()  {
	 r :=gin.Default()
	r.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username :=c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(http.StatusOK,gin.H{
			"message":"ok",
			"usernaem":username,
			"address":address,
		})
	})
	r.Run(":8081")
}
