package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3"`
	Abbreviation string             `json:"abbreviation,omitempty" bson:"abbreviation,omitempty" validate:"required,min=3"`
	Location     string             `json:"location,omitempty" bson:"location,omitempty" validate:"required,min=3"`
	TeamId       int                `json:"teamId,omitempty" bson:"teamId,omitempty" validate:"numeric"`
}
