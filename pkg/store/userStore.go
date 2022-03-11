package store

import (
	"context"
	"errors"
	"time"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"go.mongodb.org/mongo-driver/bson"
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

	count, err := collection.CountDocuments(ctx, bson.M{"email": newUser.Email})
	defer cancel()
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("this email already exists")
	}

	count, err = collection.CountDocuments(ctx, bson.M{"nickname": newUser.Nickname})
	defer cancel()
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("this nickname already exists")
	}

	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserStore) FindUserByNickname(userSigninNickname string) (*structure.UserSignin, error) {
	collection := us.db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var foundUser *structure.UserSignin

	err := collection.FindOne(ctx, bson.M{"nickname": userSigninNickname}).Decode(&foundUser)
	if err != nil {
		return nil, errors.New("could not find a user with this nickname")
	}

	return foundUser, nil
}
