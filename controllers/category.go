package controller

import (
	"example/go-crud/helper"
	"example/go-crud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

var categoryService = new(services.CategoryService)

func (ctr CategoryController) GetCategories(c *gin.Context) {
	filter := helper.ParseFilterParams(c.Request.URL.Query())
	if result, err := categoryService.GetCategories(filter); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
