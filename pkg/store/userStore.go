package store

import (
	"context"
	"time"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	db *mongo.Database
}

func NewUserStore(db *mongo.Database) *UserStore {
	return &UserStore{db: db}
}

func (us *UserStore) Signup(userSignupDTO *structure.User) error {
	collection := us.db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newUser *structure.User = userSignupDTO

	_, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}
