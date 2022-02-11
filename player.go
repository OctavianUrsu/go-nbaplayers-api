package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	PlayerId  primitive.ObjectID `json:"playerId,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	TeamId    int                `json:"teamId,omitempty" bson:"teamId,omitempty"`
}
