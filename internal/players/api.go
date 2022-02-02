package players

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type PlayersResource struct{}
type Player struct {
	PlayerId  int    `json:"playerId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	TeamId    int    `json:"teamId"`
}

const playersPath string = "./internal/players/players.json"

func (pr PlayersResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", pr.Players)
	r.Post("/", pr.Create)

	return r
}

func (pr PlayersResource) Players(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

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

	json.NewEncoder(w).Encode(allPlayers)
}

func (pr PlayersResource) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	// Declare variable for new player
	var newPlayer Player

	// Read request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal new player info to new player var
	json.Unmarshal(req, &newPlayer)

	// If request has no data log an error
	if newPlayer.FirstName == "" || newPlayer.LastName == "" || newPlayer.PlayerId == 0 {
		io.WriteString(w, "Please complete all required fields!")
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

		io.WriteString(w, "New player created!")
	}
}
