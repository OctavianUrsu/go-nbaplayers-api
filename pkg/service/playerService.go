package service

import (
	"context"
	"errors"
	"time"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/helpers"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

type PlayerService struct {
	helpers *helpers.Helpers
	storage *storage.Storage
}

// Constructor for dependency injection
func NewService(h *helpers.Helpers, r *storage.Storage) *PlayerService {
	return &PlayerService{h, r}
}

// Request Service - GET /players - Get all players.
func (ps *PlayerService) GetAll() ([]*playerStruct.Player, error) {
	collection := ps.storage.MongoDb().Collection("players")

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

// Request Service - POST /players - Add new player.
func (ps *PlayerService) Create(playerDTO playerStruct.Player) (*playerStruct.Player, error) {
	// Check if request has empty strings
	if playerDTO.FirstName == "" && playerDTO.LastName == "" {
		return &playerDTO, errors.New("complete the required fields")
	} else {
		collection := ps.storage.MongoDb().Collection("players")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var newPlayer *playerStruct.Player = &playerDTO

		_, err := collection.InsertOne(ctx, newPlayer)
		if err != nil {
			return &playerDTO, err
		}

		return newPlayer, nil
	}
}

// Request Service - GET /players/{id} - Get player by Id.
func (ps *PlayerService) GetById(id int) (*playerStruct.Player, error) {
	collection := ps.storage.MongoDb().Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var player *playerStruct.Player

	err := collection.FindOne(ctx, playerStruct.Player{PlayerId: id}).Decode(&player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// Request Service - PUT /players/{id} - Update player by Id.
func (ps *PlayerService) Update(id int, playerDTO playerStruct.Player) error {
	// Check if request has empty strings
	if playerDTO.FirstName == "" && playerDTO.LastName == "" && playerDTO.PlayerId == 0 {
		return errors.New("complete the required fields")
	} else {
		collection := ps.storage.MongoDb().Collection("players")

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
}

// Request Service - DELETE /players/{id} - Delete player by Id.
func (ps *PlayerService) Delete(id int) error {
	collection := ps.storage.MongoDb().Collection("players")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
