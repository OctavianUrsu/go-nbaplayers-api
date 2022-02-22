package store

import (
	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPlayerStore interface {
	GetAll() ([]*structure.Player, error)
	Create(playerDTO *structure.Player) error
	GetById(id string) (*structure.Player, error)
	Update(id string, playerDTO *structure.Player) error
	Delete(id string) error
	GetByName(name []string) ([]*structure.Player, error)
}

type ITeamStore interface {
	GetAllTeams() ([]*structure.Team, error)
	CreateTeam(teamDTO *structure.Team) error
	GetTeamById(id string) (*structure.Team, error)
	UpdateTeam(id string, teamDTO *structure.Team) error
	DeleteTeam(id string) error
}

type Store struct {
	IPlayerStore
	ITeamStore
}

func NewStore(db *mongo.Database) *Store {
	return &Store{
		IPlayerStore: NewPlayerStore(db),
		ITeamStore:   NewTeamStore(db),
	}
}
