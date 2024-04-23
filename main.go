package main

import (
	router "example/go-crud/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	{
		router.CategoryRouter(r.Group(""))
	}

	r.Run("localhost:8080")
}
