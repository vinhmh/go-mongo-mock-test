package services

import (
	"example/go-crud/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestGetCategories(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	// test with success case
	mt.Run("get list success", func(mt *mtest.T) {
		mockCollection := mt.Coll
		categoryService := CategoryServiceInit(mockCollection)

		// Insert test data
		first := mtest.CreateCursorResponse(1, "db.test", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "name", Value: "Category1"},
			{Key: "description", Value: "Description for category1"},
		})
		second := mtest.CreateCursorResponse(1, "db.test", mtest.NextBatch, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "name", Value: "Category2"},
			{Key: "description", Value: "Description for category2"},
		})
		killCursors := mtest.CreateCursorResponse(0, "db.test", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		filter := models.Filter{
			Limit:  20,
			Offset: 0,
			FilterDetails: []models.FilterDetail{
				{FieldName: "name", Value: "category"},
			},
		}
		response, err := categoryService.GetCategories(filter)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
	})

	// test with failed case
	mt.Run("get list failed", func(mt *mtest.T) {
		mockCollection := mt.Coll
		categoryService := CategoryServiceInit(mockCollection)
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 123, Message: "An exception occur"}))
		filter := models.Filter{
			Limit:  20,
			Offset: 0,
			FilterDetails: []models.FilterDetail{
				{FieldName: "name", Value: "category"},
			},
		}
		response, err := categoryService.GetCategories(filter)
		assert.Empty(t, response)
		assert.Error(t, err)
	})

}

func TestCreateCategory(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	var createData = models.CategoryStruct{
		Id:          primitive.NewObjectID(),
		Name:        "Category1",
		Description: "Description of category1",
	}
	// test with success case
	mt.Run("create success", func(mt *mtest.T) {
		mockCollection := mt.Coll
		categoryService := CategoryServiceInit(mockCollection)
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := categoryService.CreateCategory(createData)
		assert.NoError(t, err)
	})

	// test with failed cases
	mt.Run("create failed by duplicate key", func(mt *mtest.T) {
		mockCollection := mt.Coll
		categoryService := CategoryServiceInit(mockCollection)
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		err := categoryService.CreateCategory(createData)
		assert.NotNil(t, err)
	})

	mt.Run("create failed by missing required field", func(mt *mtest.T) {
		mockCollection := mt.Coll
		categoryService := CategoryServiceInit(mockCollection)

		createMissingRequiredFieldData := models.CategoryStruct{
			Name: "",
		}
		err := categoryService.CreateCategory(createMissingRequiredFieldData)
		assert.NotNil(t, err)
	})

	mt.Run("create failed by another reason", func(mt *mtest.T) {
		mockCollection := mt.Coll
		categoryService := CategoryServiceInit(mockCollection)
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{Code: 123, Message: "An exception occur"}))
		err := categoryService.CreateCategory(createData)
		assert.NotNil(t, err)
	})
}
