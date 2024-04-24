package main

import (
	"encoding/json"
	controller "example/go-crud/controllers"
	"example/go-crud/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetCategoriesHandler(t *testing.T) {
	r := SetUpRouter()
	categoryController := new(controller.CategoryController)
	r.GET("/categories", categoryController.GetCategories)
	req, _ := http.NewRequest("GET", "/categories", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var categories []models.CategoryStruct
	json.Unmarshal(w.Body.Bytes(), &categories)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, categories)
}
