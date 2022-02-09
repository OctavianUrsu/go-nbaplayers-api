package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/helpers"
	"github.com/sirupsen/logrus"
)

type PlayerService struct {
	helpers *helpers.Helpers
}

// Path to players JSON
const playersJsonPath = "./players.json"

// Constructor for dependency injection
func NewService(h *helpers.Helpers) *PlayerService {
	return &PlayerService{h}
}

// Request Service - GET /players - Get all players.
func (ps *PlayerService) GetAll() []playerStruct.Player {
	// Get all players
	allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

	return allPlayers
}

// Request Service - POST /players - Add new player.
func (ps *PlayerService) Create(playerDTO playerStruct.Player) (playerStruct.Player, error) {
	// Check if request has empty strings
	if playerDTO.FirstName == "" || playerDTO.LastName == "" || playerDTO.PlayerId == 0 {
		return playerDTO, errors.New("complete the required fields")
	} else {
		// Get all players
		allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

		// Check if player with request id already exists
		for _, player := range allPlayers {
			if player.PlayerId == playerDTO.PlayerId {
				return playerDTO, errors.New("a player with this id already exists")
			}
		}

		// Create a variable to store info about new player
		var newPlayer playerStruct.Player = playerDTO

		// Add new player to all players
		allPlayers = append(allPlayers, newPlayer)

		// Encode all players in JSON format
		byteValue, err := json.Marshal(allPlayers)
		if err != nil {
			logrus.New().Warn(err)
		}

		// Write the encoded JSON to players JSON file
		if err = ioutil.WriteFile(playersJsonPath, byteValue, 0644); err != nil {
			logrus.New().Warn(err)
		}

		return newPlayer, nil
	}
}

// Request Service - GET /players/{id} - Get player by Id.
func (ps *PlayerService) GetById(id int) playerStruct.Player {
	// Get all players
	allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

	// Create a variable to store our player specific to id
	var player playerStruct.Player

	// Loop through all players and retrieve info about player specific to id
	for _, p := range allPlayers {
		if p.PlayerId == id {
			player = p
		}
	}

	return player
}

// Request Service - PUT /players/{id} - Update player by Id.
func (ps *PlayerService) Update(id int, playerDTO playerStruct.Player) (playerStruct.Player, error) {
	// Check if request has empty strings
	if playerDTO.FirstName == "" || playerDTO.LastName == "" {
		return playerDTO, errors.New("complete the required fields")
	} else {
		// Get all players
		allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

		// Create a variable that stores info about player update
		var updatePlayer playerStruct.Player = playerDTO

		// Assign id from request to updated player
		updatePlayer.PlayerId = id

		// Loop all players
		for i, p := range allPlayers {
			// Check if id is the same from request
			if p.PlayerId == id {
				// Update the player information
				p.FirstName = updatePlayer.FirstName
				p.LastName = updatePlayer.LastName
				p.TeamId = updatePlayer.TeamId
				allPlayers = append(allPlayers[:i], p)
			}
		}

		// Encode all players in JSON format
		byteValue, err := json.Marshal(allPlayers)
		if err != nil {
			logrus.New().Warn(err)
		}

		// Write the encoded JSON to players JSON file
		err = ioutil.WriteFile(playersJsonPath, byteValue, 0644)
		if err != nil {
			logrus.New().Warn(err)
		}

		return updatePlayer, nil
	}
}

// Request Service - DELETE /players/{id} - Delete player by Id.
func (ps *PlayerService) Delete(id int) playerStruct.Player {
	// Get all players
	allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

	// Create a variable to store our deleted player info
	var playerToBeDeleted playerStruct.Player

	// Loop all players
	for i, p := range allPlayers {
		// Check if id is the same from request
		if p.PlayerId == id {
			// Delete player
			allPlayers = append(allPlayers[:i], allPlayers[i+1:]...)
			playerToBeDeleted = p
		}
	}

	// Encode all players in JSON format
	byteValue, err := json.Marshal(allPlayers)
	if err != nil {
		logrus.New().Warn(err)
	}

	// Write the encoded JSON to players JSON file
	err = ioutil.WriteFile(playersJsonPath, byteValue, 0644)
	if err != nil {
		logrus.New().Warn(err)
	}

	return playerToBeDeleted
}
