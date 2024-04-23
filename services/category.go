package services

import (
	"context"
	"example/go-crud/config"
	"example/go-crud/forms"
	"example/go-crud/helper"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryService struct{}

var categoryCollection = config.GetCollection("categories")

func (s CategoryService) GetCategories(filter *forms.Filter) (response []forms.CategoryStruct, err error) {
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
	cursor, err := categoryCollection.Find(context.Background(), qr, opts)
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.Background(), &response)
	return
}
