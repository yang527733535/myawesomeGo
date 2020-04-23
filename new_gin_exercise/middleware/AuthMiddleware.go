package middleware

import (
	"github.com/gin-gonic/gin"
	"mymod/common"
	"mymod/model"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 TOKEN
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParaseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		//验证通过后 获取Userid
		userId := claims.UserId
		DB := common.GetDb()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		//用户存在 将User信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
