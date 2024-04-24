package services

import (
	"context"
	"example/go-crud/helper"
	"example/go-crud/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryService struct{ Collection *mongo.Collection }

func CategoryServiceInit(collection *mongo.Collection) *CategoryService {
	return &CategoryService{
		Collection: collection,
	}
}

func (s CategoryService) GetCategories(filter models.Filter) (response []models.CategoryStruct, err error) {
	qr := bson.M{}
	opts := options.Find()
	opts.SetLimit(filter.Limit)
	opts.SetSkip(filter.Offset)
	if len(filter.FilterDetails) > 0 {
		for _, k := range filter.FilterDetails {
			if k.FieldName != "" && k.Value != nil {
				qr = helper.AppendCondition(qr, k.FieldName, bson.M{"$regex": k.Value})
			}
		}
	}
	cursor, err := s.Collection.Find(context.Background(), qr, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &response)
	return
}

func (s CategoryService) CreateCategory(createData models.CategoryStruct) error {
	if createData.Name == "" {
		return fmt.Errorf("Category name is require")
	}
	_, err := s.Collection.InsertOne(context.Background(), createData)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("Category name %s is duplicated!", createData.Name)
		}
		return err
	}
	return nil
}
