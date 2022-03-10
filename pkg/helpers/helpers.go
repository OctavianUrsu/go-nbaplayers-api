package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Helpers struct{}

// Constructor for dependency injection
func NewHelper() *Helpers {
	return &Helpers{}
}

func (h *Helpers) UnmarshalPlayersJson(path string) []playerStruct.Player {
	// Read all players from the JSON file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	// Declare a variable to store all players
	var allPlayers []playerStruct.Player

	// Unmarshal the JSON file to allPlayers var
	json.Unmarshal(file, &allPlayers)

	return allPlayers
}

func (h *Helpers) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		logrus.Panic(err)
	}

	return string(bytes)
}

func (h *Helpers) VerifyPassword(userPassword string, providedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	if err != nil {
		return false, err
	}

	return true, nil
}
