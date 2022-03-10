package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name" validate:"required,min=3"`
	Abbreviation string             `json:"abbreviation" bson:"abbreviation" validate:"required,min=3"`
	Location     string             `json:"location" bson:"location" validate:"required,min=3"`
	TeamId       int                `json:"teamId,omitempty" bson:"teamId,omitempty"`
}
