package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Collection struct {
	Mongo_id     primitive.ObjectID `json:"mongo_id,omitempty"`
	Id           string             `json:"id,omitempty" validate:"required"`
	Stac_version string             `json:"stac_version,omitempty" validate:"required"`
	Description  string             `json:"description,omitempty" validate:"required"`
}
