package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct{}

// Constructor for dependency injection
func NewStorage() *Storage {
	return &Storage{}
}

func (r *Storage) MongoDb() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://octavian:4FQu*yym2V@vault.bkm5p.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	db := client.Database("nba")

	return db
}
