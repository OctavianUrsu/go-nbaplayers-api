package players

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const playersPath string = "./internal/players/players.json"

type Player struct {
	PlayerId  int    `json:"playerId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	TeamId    int    `json:"teamId"`
}

func HandlePlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		allPlayers := getAllPlayers()

		// Encode our array into the response
		json.NewEncoder(w).Encode(allPlayers)

	case "POST":
		w.WriteHeader(http.StatusCreated)

		req, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		postRes := createNewPlayer(req)
		io.WriteString(w, postRes)
	}
}

func getAllPlayers() []Player {
	// Open JSON file and if err => handle it
	playersJson, err := os.Open(playersPath)
	if err != nil {
		fmt.Println(err)
	}

	// Defer closing the file
	defer playersJson.Close()

	// Read the file as a byte array
	byteValue, _ := ioutil.ReadAll(playersJson)

	// Initialize the array that will store players
	var allPlayers []Player

	// Unmarshal the byteValue that contains our json info
	// and store it in our array
	json.Unmarshal(byteValue, &allPlayers)

	return allPlayers
}

func createNewPlayer(p []byte) string {
	// Declare variable for new player
	var newPlayer Player

	// Unmarshal new player info to new player var
	json.Unmarshal(p, &newPlayer)

	// If request has no data log an error
	if newPlayer.FirstName == "" || newPlayer.LastName == "" || newPlayer.PlayerId == 0 {
		return "Please complete all required fields!"
	} else {
		// Read all players from the JSON fileDB
		file, err := ioutil.ReadFile(playersPath)
		if err != nil {
			fmt.Println(err)
		}

		// Declare a variable to store all players
		var allPlayers []Player

		// Unmarshal the JSON fileDB to allPlayers var
		json.Unmarshal(file, &allPlayers)

		// Append the new player to all players
		allPlayers = append(allPlayers, newPlayer)

		// Marshal back all players into JSON
		byteValue, err := json.Marshal(allPlayers)
		if err != nil {
			fmt.Println(err)
		}

		// Write the new all players JSON to the fileDB
		err = ioutil.WriteFile(playersPath, byteValue, 0644)
		if err != nil {
			fmt.Println(err)
		}

		return "New player created!"
	}
}
