package store

import (
	"context"
	"errors"
	"time"

	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamStore struct {
	db *mongo.Database
}

func NewTeamStore(db *mongo.Database) *TeamStore {
	return &TeamStore{db: db}
}

func (ts *TeamStore) GetAllTeams() ([]*models.Team, error) {
	collection := ts.db.Collection("teams")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var allTeams []*models.Team

	for cursor.Next(ctx) {
		team := &models.Team{}
		err := cursor.Decode(team)
		if err != nil {
			return nil, err
		}

		allTeams = append(allTeams, team)
	}

	return allTeams, nil
}

func (ts *TeamStore) CreateTeam(teamDTO *models.Team) error {
	collection := ts.db.Collection("teams")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newTeam *models.Team = teamDTO

	_, err := collection.InsertOne(ctx, newTeam)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TeamStore) GetTeamById(id string) (*models.Team, error) {
	collection := ts.db.Collection("teams")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var team *models.Team

	objId, _ := primitive.ObjectIDFromHex(id)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (ts *TeamStore) UpdateTeam(id string, teamDTO *models.Team) error {
	collection := ts.db.Collection("teams")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a variable that stores info about team update
	var updateTeam = bson.M{
		"$set": teamDTO,
	}

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objId}, updateTeam)
	if result.ModifiedCount == 0 {
		return errors.New("could not find a team with this id")
	}

	if err != nil {
		return err
	}

	return nil
}

func (ts *TeamStore) DeleteTeam(id string) error {
	collection := ts.db.Collection("teams")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objId})
	if result.DeletedCount == 0 {
		return errors.New("could not find a team with this id")
	}

	if err != nil {
		return err
	}

	return nil
}
