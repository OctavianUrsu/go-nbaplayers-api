package api

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nickname      string             `json:"nickname" validate:"required,min=3,max=12"`
	Email         string             `json:"email" validate:"email,required"`
	Password      string             `json:"password" validate:"required,min=6"`
	Token         string             `json:"token,omitempty"`
	Refresh_token string             `json:"refresh_token,omitempty"`
	Created_at    time.Time          `json:"created_at,omitempty"`
	Updated_at    time.Time          `json:"updated_at,omitempty"`
}
