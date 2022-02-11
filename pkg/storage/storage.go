package storage

import (
	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPlayerStorage interface {
	GetAll() ([]*playerStruct.Player, error)
	Create(playerDTO *playerStruct.Player) error
	GetById(id string) (*playerStruct.Player, error)
	Update(id string, playerDTO *playerStruct.Player) error
	Delete(id string) error
}

type Storage struct {
	IPlayerStorage
}

func NewStorage(db *mongo.Database) *Storage {
	return &Storage{
		IPlayerStorage: NewPlayerStorage(db),
	}
}
