package store

import (
	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPlayerStore interface {
	GetAllPlayers() ([]*structure.Player, error)
	CreatePlayer(playerDTO *structure.Player) error
	GetPlayerById(id string) (*structure.Player, error)
	UpdatePlayer(id string, playerDTO *structure.Player) error
	DeletePlayer(id string) error
	GetPlayerByName(name []string) ([]*structure.Player, error)
}

type ITeamStore interface {
	GetAllTeams() ([]*structure.Team, error)
	CreateTeam(teamDTO *structure.Team) error
	GetTeamById(id string) (*structure.Team, error)
	UpdateTeam(id string, teamDTO *structure.Team) error
	DeleteTeam(id string) error
}

type IUserStore interface {
	Signup(userSignupDTO *structure.User) error
	FindUserByNickname(userSigninNickname string) (*structure.UserSignin, error)
	FindUserByTokenClaims(claims *structure.SignedClaims) (*bool, error)
}

type Store struct {
	IPlayerStore
	ITeamStore
	IUserStore
}

func NewStore(db *mongo.Database) *Store {
	return &Store{
		IPlayerStore: NewPlayerStore(db),
		ITeamStore:   NewTeamStore(db),
		IUserStore:   NewUserStore(db),
	}
}
