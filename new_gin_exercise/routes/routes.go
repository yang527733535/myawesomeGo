package routes

import (
	"mymod/controller"
	"mymod/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/User/EditName", controller.EditUserName)
	r.DELETE("/api/User/delete/:id", controller.Delete)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	r.GET("/api/userlist", controller.Requserlist)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT(":id", categoryController.Update)
	categoryRoutes.GET(":id", categoryController.Show)
	categoryRoutes.DELETE(":id", categoryController.Delete)

	//UserRoutes := r.Group("/categories")
	//categoryRoutes.PUT(":id", categoryController.Update)
	return r
}
