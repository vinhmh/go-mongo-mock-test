package main

import (
	"example/go-crud/docs"
	router "example/go-crud/routers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}
	docs.SwaggerInfo.BasePath = ""
	r := gin.Default()
	{
		router.CategoryRouter(r.Group(""))
	}

	r.Static("/public", "./public")

	// Define a route to serve the HTML file
	r.GET("/public", func(c *gin.Context) {
		c.HTML(http.StatusOK, "coverage.html", nil)
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
