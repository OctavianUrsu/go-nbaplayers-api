package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
)

type PlayerService struct{}

const playersJsonPath = "./players.json"

// Request Service - GET /players - Get all players.
func (ps *PlayerService) GetAll() []playerStruct.Player {
	// Open JSON file and if err => handle it
	playersJson, err := os.Open(playersJsonPath)
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
	// Open JSON file and if err => handle it
	playersJson, err := os.Open(playersJsonPath)
	if err != nil {
		fmt.Println(err)
	}

	// Defer closing the file
	defer playersJson.Close()

	// Read the file as a byte array
	byteValue, _ := ioutil.ReadAll(playersJson)

	// Initialize the array that will store players
	var allPlayers []playerStruct.Player
	var newPlayer playerStruct.Player

	// Unmarshal the byteValue that contains our json info
	// and store it in our array
	json.Unmarshal(byteValue, &allPlayers)

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
	// Open JSON file and if err => handle it
	playersJson, err := os.Open(playersJsonPath)
	if err != nil {
		fmt.Println(err)
	}

	// Defer closing the file
	defer playersJson.Close()

	// Read the file as a byte array
	byteValue, _ := ioutil.ReadAll(playersJson)

	// Initialize the array that will store players
	var allPlayers []playerStruct.Player
	var updatePlayer playerStruct.Player = playerDTO
	updatePlayer.PlayerId = id

	// Unmarshal the byteValue that contains our json info
	// and store it in our array
	json.Unmarshal(byteValue, &allPlayers)

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
	byteValue, err = json.Marshal(allPlayers)
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
	// Open JSON file and if err => handle it
	playersJson, err := os.Open(playersJsonPath)
	if err != nil {
		fmt.Println(err)
	}

	// Defer closing the file
	defer playersJson.Close()

	// Read the file as a byte array
	byteValue, _ := ioutil.ReadAll(playersJson)

	// Initialize the array that will store players
	var allPlayers []playerStruct.Player
	var deletedPlayer playerStruct.Player

	// Unmarshal the byteValue that contains our json info
	// and store it in our array
	json.Unmarshal(byteValue, &allPlayers)

	// Loop through all players and
	// find the one that matches our id
	for i, player := range allPlayers {
		if player.PlayerId == id {
			allPlayers = append(allPlayers[:i], allPlayers[i+1:]...)
			deletedPlayer = player
		}
	}

	// Marshal back all players into JSON
	byteValue, err = json.Marshal(allPlayers)
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
