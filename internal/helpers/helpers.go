package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
)

type Helpers struct{}

func (h *Helpers) UnmarshalPlayersJson(path string) []playerStruct.Player {
	// Read all players from the JSON fileDB
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	// Declare a variable to store all players
	var allPlayers []playerStruct.Player

	// Unmarshal the JSON fileDB to allPlayers var
	json.Unmarshal(file, &allPlayers)

	return allPlayers
}
