package controller

import (
	"example/go-crud/connections"
	"example/go-crud/helper"
	"example/go-crud/models"
	"example/go-crud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

var collection = connections.GetCollection("categories")
var categoryService = services.CategoryServiceInit(collection)

// @Summary Get categories
// @Description Get all categories
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int false "Offset"
// @Param filterDetails query string false "Filter details" Example("[{\"fieldName\":\"name\",\"value\":\"category1\"}]")
// @Success 200 {array} array
// @Router /categories [get]
func (ctr CategoryController) GetCategories(c *gin.Context) {
	filter := helper.ParseFilterParams(c.Request.URL.Query())
	if result, err := categoryService.GetCategories(filter); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// @Summary Create a category
// @Description Create a new category
// @Accept json
// @Produce json
// @Param categoryData body models.CategoryStruct true "Category data"
// @Success 200 {object} nil
// @Failure	400
// @Router /categories [post]
func (ctr CategoryController) CreateCategory(c *gin.Context) {
	createData := models.CategoryStruct{}
	err := c.Bind(&createData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = categoryService.CreateCategory(createData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}
