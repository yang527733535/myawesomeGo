package routes

import (
	"github.com/gin-gonic/gin"
	"mymod/controller"
	"mymod/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	r.GET("/api/userlist", controller.Requserlist)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()

	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT(":id", categoryController.Update)
	categoryRoutes.GET(":id", categoryController.Show)
	categoryRoutes.DELETE(":id", categoryController.Delete)
	return r
}
