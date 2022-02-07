package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/helpers"
)

type PlayerService struct {
	helpers *helpers.Helpers
}

// JSON Path
const playersJsonPath = "./players.json"

// Constructor
func NewService(h *helpers.Helpers) *PlayerService {
	return &PlayerService{h}
}

// Request Service - GET /players - Get all players.
func (ps *PlayerService) GetAll() []playerStruct.Player {
	allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

	return allPlayers
}

// Request Service - POST /players - Add new player.
func (ps *PlayerService) Create(playerDTO playerStruct.Player) (playerStruct.Player, error) {
	if playerDTO.FirstName == "" || playerDTO.LastName == "" || playerDTO.PlayerId == 0 {
		return playerDTO, errors.New("complete the required fields")
	} else {
		// Read all players from the JSON fileDB
		file, err := ioutil.ReadFile(playersJsonPath)
		if err != nil {
			fmt.Println(err)
		}

		// Declare a variable to store all players
		var allPlayers []playerStruct.Player

		// Unmarshal the JSON fileDB to allPlayers var
		json.Unmarshal(file, &allPlayers)

		// Append the new player to all players
		allPlayers = append(allPlayers, playerDTO)

		// Marshal back all players into JSON
		byteValue, err := json.Marshal(allPlayers)
		if err != nil {
			fmt.Println(err)
		}

		// Write the new all players JSON to the fileDB
		err = ioutil.WriteFile(playersJsonPath, byteValue, 0644)
		if err != nil {
			fmt.Println(err)
		}

		return playerDTO, nil

	}
}

// Request Service - GET /players/{id} - Get player by Id.
func (ps *PlayerService) GetById(id int) playerStruct.Player {
	allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

	// Initialize the array that will store players
	var newPlayer playerStruct.Player

	// Loop through all players and
	// find the one that matches our id
	for _, player := range allPlayers {
		if player.PlayerId == id {
			newPlayer = player
		}
	}

	return newPlayer
}

// Request Service - PUT /players/{id} - Update player by Id.
func (ps *PlayerService) Update(id int, playerDTO playerStruct.Player) playerStruct.Player {
	allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

	// Initialize the array that will store players
	var updatePlayer playerStruct.Player = playerDTO
	updatePlayer.PlayerId = id

	// Loop through all players and
	// find the one that matches our id
	for i, player := range allPlayers {
		if player.PlayerId == id {
			player.FirstName = playerDTO.FirstName
			player.LastName = playerDTO.LastName
			player.TeamId = playerDTO.TeamId
			allPlayers = append(allPlayers[:i], player)
		}
	}

	// Marshal back all players into JSON
	byteValue, err := json.Marshal(allPlayers)
	if err != nil {
		fmt.Println(err)
	}

	// Write the new all players JSON to the fileDB
	err = ioutil.WriteFile(playersJsonPath, byteValue, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return updatePlayer
}

// Request Service - DELETE /players/{id} - Delete player by Id.
func (ps *PlayerService) Delete(id int) playerStruct.Player {
	allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

	// Initialize the array that will store players
	var deletedPlayer playerStruct.Player

	// Loop through all players and
	// find the one that matches our id
	for i, player := range allPlayers {
		if player.PlayerId == id {
			allPlayers = append(allPlayers[:i], allPlayers[i+1:]...)
			deletedPlayer = player
		}
	}

	// Marshal back all players into JSON
	byteValue, err := json.Marshal(allPlayers)
	if err != nil {
		fmt.Println(err)
	}

	// Write the new all players JSON to the fileDB
	err = ioutil.WriteFile(playersJsonPath, byteValue, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return deletedPlayer
}
