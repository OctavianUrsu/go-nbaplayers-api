package storage

import (
	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPlayerStorage interface {
	GetAll() ([]*playerStruct.Player, error)
	Create(playerDTO *playerStruct.Player) error
	GetById(id int) (*playerStruct.Player, error)
	Update(id int, playerDTO *playerStruct.Player) error
	Delete(id int) error
}

type Storage struct {
	IPlayerStorage
}

func NewStorage(db *mongo.Database) *Storage {
	return &Storage{
		IPlayerStorage: NewPlayerStorage(db),
	}
}
