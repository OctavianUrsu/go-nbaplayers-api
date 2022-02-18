package store

import (
	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPlayerStore interface {
	GetAll() ([]*playerStruct.Player, error)
	Create(playerDTO *playerStruct.Player) error
	GetById(id string) (*playerStruct.Player, error)
	Update(id string, playerDTO *playerStruct.Player) error
	Delete(id string) error
	GetByName(name []string) ([]*playerStruct.Player, error)
}

type Store struct {
	IPlayerStore
}

func NewStore(db *mongo.Database) *Store {
	return &Store{
		IPlayerStore: NewPlayerStore(db),
	}
}
