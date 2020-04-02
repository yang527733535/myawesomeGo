package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main()  {
	r:=gin.Default()
	// 注册一个全局中间件
	//r.Use(StatCost())
	r.GET("/test",StatCost(), func(context *gin.Context) {
		name:=context.MustGet("name").(string)
		cost:=context.MustGet("cost")
		log.Println(name)
		log.Println(cost)
			context.JSON(200,gin.H{
			"name":name,
		})
	})
	r.Run(":8081")
}


// StatCost 是一个统计耗时请求耗时的中间件
//Gin中的中间件必须是一个gin.HandlerFunc类型。例如我们像下面的代码一样定义一个统计请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
	start:=time.Now()
	c.Set("name","yang")
		cost := time.Since(start)
		c.Set("cost",cost)
		log.Println(cost)
	    // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()


	}
}