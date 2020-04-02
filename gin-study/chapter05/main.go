package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//参数绑定
type Login struct {
	User string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main()  {
	router := gin.Default()

	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	router.POST("/loginJSON", func(context *gin.Context) {
		var login Login
			err:=context.ShouldBind(&login)
			if err==nil{
				fmt.Printf("login info:%#v\n", login)
				context.JSON(200,gin.H{
					"user":login.User,
					"password":login.Password,
				})
			}else{
				context.JSON(http.StatusBadRequest,gin.H{
					"error":err.Error(),
				})
			}
	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	router.POST("/loginForm", func(context *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
	 err:=context.ShouldBind(&login);
	 if err==nil{
	 	context.JSON(http.StatusOK,gin.H{
	 		"user":login.User,
	 		"password":login.Password,
		})
	 }else{
	context.JSON(http.StatusBadRequest,gin.H{
		"error":err.Error(),
	})
	 }
	})

	router.GET("loginForm", func(context *gin.Context) {

		var login Login
		err:=context.ShouldBind(&login)
		if err==nil{
			//time.Sleep(time.Second*5)
			context.JSON(200,gin.H{
				"user":login.User,
				"pwd":login.Password,
			})
		}else{
			context.JSON(400,gin.H{
				"error":err.Error(),
			})
		}

	})
	router.Run(":8081")
}

//
//ShouldBind会按照下面的顺序解析请求中的数据完成绑定：
//
//如果是 GET 请求，只使用 Form 绑定引擎（query）。
//如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）。