package store

import (
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPlayerStore interface {
	GetAllPlayers() ([]*models.Player, error)
	CreatePlayer(playerDTO *models.Player) error
	GetPlayerById(id string) (*models.Player, error)
	UpdatePlayer(id string, playerDTO *models.Player) error
	DeletePlayer(id string) error
	GetPlayerByName(name []string) ([]*models.Player, error)
}

type ITeamStore interface {
	GetAllTeams() ([]*models.Team, error)
	CreateTeam(teamDTO *models.Team) error
	GetTeamById(id string) (*models.Team, error)
	UpdateTeam(id string, teamDTO *models.Team) error
	DeleteTeam(id string) error
}

type IUserStore interface {
	Signup(userSignupDTO *models.User) error
	NicknameExists(user *models.User) bool
	EmailExists(user *models.User) bool
	GetHashedPassword(userSigninNickname string) (string, error)
	//FindUserByTokenClaims(claims *models.SignedClaims) (*bool, error)
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
