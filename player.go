package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	PlayerId  primitive.ObjectID `json:"playerId,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName" validate:"required,min=3"`
	LastName  string             `json:"lastName" bson:"lastName" validate:"required,min=3"`
	TeamId    int                `json:"teamId,omitempty" bson:"teamId,omitempty"`
}
