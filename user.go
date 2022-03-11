package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nickname   string             `json:"nickname" validate:"required,min=3,max=12"`
	Email      string             `json:"email" validate:"email,required"`
	Password   string             `json:"password" validate:"required,min=6"`
	Created_at time.Time          `json:"created_at,omitempty"`
}

type UserSignin struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type SignedClaims struct {
	Nickname string
	Password string
	jwt.StandardClaims
}
