package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI        string
	Name       string
	Password   string
	Collection string
}

func NewMongoDB(cfg Config) (*mongo.Database, error) {
	var databaseURI string = fmt.Sprintf("mongodb+srv://%s:%s@%s/myFirstDatabase?retryWrites=true&w=majority", cfg.Name, cfg.Password, cfg.URI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(databaseURI))
	db := client.Database("nba")

	return db, nil
}
