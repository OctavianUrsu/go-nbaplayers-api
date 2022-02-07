package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
)

type Helpers struct{}

func (h *Helpers) UnmarshalPlayersJson(path string) []playerStruct.Player {
	fmt.Println("running helper function")

	// Open JSON file and if err => handle it
	playersJson, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	// Defer closing the file
	defer playersJson.Close()

	// Read the file as a byte array
	byteValue, _ := ioutil.ReadAll(playersJson)

	// Initialize the array that will store players
	var allPlayers []playerStruct.Player

	// Unmarshal the byteValue that contains our json info
	// and store it in our array
	json.Unmarshal(byteValue, &allPlayers)

	return allPlayers
}
