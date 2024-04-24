package router

import (
	controller "example/go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(route *gin.RouterGroup) {
	categoryRouter := route.Group("/categories")
	{
		categoryController := new(controller.CategoryController)
		categoryRouter.GET("", categoryController.GetCategories)
		categoryRouter.POST("", categoryController.CreateCategory)
		// categoryRouter.PUT("/change-password", middlewares.TokenAuthMiddleware(), categoryController.ChangePassword)
	}
}
