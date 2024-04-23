package controller

import (
	"example/go-crud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

var categoryService = new(services.CategoryService)

func (ctr CategoryController) GetCategories(c *gin.Context) {
	// filter := parseFilterParams(r.URL.Query())
	c.Request
	if result, err := categoryService.GetCategories(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.DefaultResponse(false, nil, err.Error()))
	} else {
		c.JSON(http.StatusOK, response.DefaultResponse(true, result, ""))
	}
}
