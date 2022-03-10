package store

import (
	"context"
	"errors"
	"time"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerStore struct {
	db *mongo.Database
}

func NewPlayerStore(db *mongo.Database) *PlayerStore {
	return &PlayerStore{db: db}
}

func (ps *PlayerStore) GetAllPlayers() ([]*structure.Player, error) {
	collection := ps.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var allPlayers []*structure.Player

	for cursor.Next(ctx) {
		player := &structure.Player{}
		err := cursor.Decode(player)
		if err != nil {
			return nil, err
		}

		allPlayers = append(allPlayers, player)
	}

	return allPlayers, nil
}

func (ps *PlayerStore) CreatePlayer(playerDTO *structure.Player) error {
	collection := ps.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newPlayer *structure.Player = playerDTO

	_, err := collection.InsertOne(ctx, newPlayer)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PlayerStore) GetPlayerById(id string) (*structure.Player, error) {
	collection := ps.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var player *structure.Player

	objId, _ := primitive.ObjectIDFromHex(id)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (ps *PlayerStore) UpdatePlayer(id string, playerDTO *structure.Player) error {
	collection := ps.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a variable that stores info about player update
	var updatePlayer = bson.M{
		"$set": playerDTO,
	}

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objId}, updatePlayer)
	if result.MatchedCount == 0 {
		return errors.New("the player with this id was not found")
	}

	if result.ModifiedCount == 0 {
		return errors.New("could not update the player")
	}

	if err != nil {
		return err
	}

	return nil
}

func (ps *PlayerStore) DeletePlayer(id string) error {
	collection := ps.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objId})
	if result.DeletedCount == 0 {
		return errors.New("could not find a user with this id")
	}

	if err != nil {
		return err
	}

	return nil
}

func (ps *PlayerStore) GetPlayerByName(searchParams []string) ([]*structure.Player, error) {
	collection := ps.db.Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := bson.M{
		"$or": []bson.M{
			{"firstName": bson.M{"$regex": searchParams[0], "$options": "i"}},
			{"lastName": bson.M{"$regex": searchParams[0], "$options": "i"}},
		}}

	if len(searchParams) > 1 {
		query = bson.M{
			"$or": []bson.M{
				{"$and": []bson.M{
					{"firstName": bson.M{"$regex": searchParams[0], "$options": "i"}},
					{"lastName": bson.M{"$regex": searchParams[1], "$options": "i"}},
				}},
				{"$and": []bson.M{
					{"firstName": bson.M{"$regex": searchParams[1], "$options": "i"}},
					{"lastName": bson.M{"$regex": searchParams[0], "$options": "i"}},
				}},
			},
		}
	}

	cursor, err := collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var foundPlayers []*structure.Player

	for cursor.Next(ctx) {
		player := &structure.Player{}
		err := cursor.Decode(player)
		if err != nil {
			return nil, err
		}

		foundPlayers = append(foundPlayers, player)
	}

	if foundPlayers == nil {
		foundPlayers = make([]*structure.Player, 0)
		return foundPlayers, nil
	}

	return foundPlayers, nil
}
