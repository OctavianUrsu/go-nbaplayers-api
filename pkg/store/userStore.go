package store

import (
	"context"
	"errors"
	"time"

	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	db *mongo.Database
}

func NewUserStore(db *mongo.Database) *UserStore {
	return &UserStore{db: db}
}

func (us *UserStore) Signup(userSignupDTO *models.User) error {
	collection := us.db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newUser *models.User = userSignupDTO

	_, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserStore) NicknameExists(user *models.User) bool {
	collection := us.db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{"nickname": user.Nickname})
	if err != nil {
		panic(err)
	}
	if count >= 1 {
		return true
	}

	return false
}

func (us *UserStore) EmailExists(user *models.User) bool {
	collection := us.db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		panic(err)
	}
	if count >= 1 {
		return true
	}

	return false
}

func (us *UserStore) GetHashedPassword(userSigninNickname string) (string, error) {
	collection := us.db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var foundUser *models.UserSignin

	err := collection.FindOne(ctx, bson.M{"nickname": userSigninNickname}).Decode(&foundUser)
	if err != nil {
		return "", errors.New("could not find a user with this nickname")
	}

	return foundUser.Password, nil
}

// func (us *UserStore) FindUserByTokenClaims(claims *models.SignedClaims) (*bool, error) {
// 	collection := us.db.Collection("users")

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	var isAuthorized bool = false

// 	query := bson.M{
// 		"$and": []bson.M{
// 			{"nickname": claims.Nickname},
// 			{"password": claims.Password},
// 		}}

// 	if err := collection.FindOne(ctx, query); err != nil {
// 		isAuthorized = true
// 	} else {
// 		return &isAuthorized, nil
// 	}

// 	return &isAuthorized, nil
// }
