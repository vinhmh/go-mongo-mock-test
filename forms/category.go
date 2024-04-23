package forms

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryStruct struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
}