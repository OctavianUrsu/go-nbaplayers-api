package storage

import (
	"context"
	"errors"
	"time"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerStorage struct {
	db *mongo.Database
}

func NewPlayerStorage(db *mongo.Database) *PlayerStorage {
	return &PlayerStorage{db: db}
}

func (pstrg *PlayerStorage) GetAll() ([]*playerStruct.Player, error) {
	collection := pstrg.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, playerStruct.Player{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var allPlayers []*playerStruct.Player

	for cursor.Next(ctx) {
		player := &playerStruct.Player{}
		err := cursor.Decode(player)
		if err != nil {
			return nil, err
		}

		allPlayers = append(allPlayers, player)
	}

	return allPlayers, nil
}

func (pstrg *PlayerStorage) Create(playerDTO *playerStruct.Player) error {
	collection := pstrg.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newPlayer *playerStruct.Player = playerDTO

	_, err := collection.InsertOne(ctx, newPlayer)
	if err != nil {
		return err
	}

	return nil
}

func (pstrg *PlayerStorage) GetById(id int) (*playerStruct.Player, error) {
	collection := pstrg.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var player *playerStruct.Player

	err := collection.FindOne(ctx, playerStruct.Player{PlayerId: id}).Decode(&player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (pstrg *PlayerStorage) Update(id int, playerDTO *playerStruct.Player) error {
	collection := pstrg.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a variable that stores info about player update
	var updatePlayer = bson.M{
		"$set": playerDTO,
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, updatePlayer)
	if err != nil {
		return err
	}

	return nil
}

func (pstrg *PlayerStorage) Delete(id int) error {
	collection := pstrg.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if result.DeletedCount == 0 {
		return errors.New("could not find a user with this id")
	}

	if err != nil {
		return err
	}

	return nil
}
