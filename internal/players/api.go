package players

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Player struct {
	PlayerId  int    `json:"playerId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	TeamId    int    `json:"teamId"`
}

func ReturnAllPlayers(w http.ResponseWriter, r *http.Request) {
	// Open JSON file and if err => handle it
	playersJson, err := os.Open("./internal/players/players.json")
	if err != nil {
		fmt.Println(err)
	}

	// Defer closing the file
	defer playersJson.Close()

	// Read the file as a byte array
	byteValue, _ := io.ReadAll(playersJson)

	// Initialize the array that will store players
	var allPlayers []Player

	// Unmarshal the byteValue that contains our json info
	// and store it in our array
	json.Unmarshal(byteValue, &allPlayers)

	// Encode our array into the response
	json.NewEncoder(w).Encode(allPlayers)

}
