package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	PlayerId  primitive.ObjectID `json:"playerId,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"firstName,omitempty" validate:"required,min=3"`
	LastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty" validate:"required,min=3"`
	TeamId    int                `json:"teamId,omitempty" bson:"teamId,omitempty" validate:"numeric"`
}
