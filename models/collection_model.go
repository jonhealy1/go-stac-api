package models

// type Link struct {
// 	Rel string `json:"rel,omitempty"`
// }

// type Collection struct {
// 	Mongo_id     primitive.ObjectID `json:"mongo_id,omitempty"`
// 	Id           string             `json:"id,omitempty" validate:"required"`
// 	Stac_version string             `json:"stac_version,omitempty" validate:"required"`
// 	Description  string             `json:"description,omitempty" validate:"required"`
// 	Links        Link               `bson:"inline"`
// }

type Collection struct {
	Id          string      `json:"id,omitempty"`
	Stac_object interface{} `json:"stac_object"`
}
