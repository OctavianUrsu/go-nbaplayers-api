package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Abbreviation string             `json:"abbreviation,omitempty" bson:"abbreviation,omitempty"`
	Location     string             `json:"location,omitempty" bson:"location,omitempty"`
	TeamId       int                `json:"teamId,omitempty" bson:"teamId,omitempty"`
}
