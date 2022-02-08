package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/helpers"
	"github.com/sirupsen/logrus"
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
		allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

		for _, player := range allPlayers {
			if player.PlayerId == playerDTO.PlayerId {
				return playerDTO, errors.New("a player with this id already exists")
			}
		}

		allPlayers = append(allPlayers, playerDTO)

		byteValue, err := json.Marshal(allPlayers)
		if err != nil {
			logrus.New().Warn(err)
		}

		if err = ioutil.WriteFile(playersJsonPath, byteValue, 0644); err != nil {
			logrus.New().Warn(err)
		}

		return playerDTO, nil
	}
}

// Request Service - GET /players/{id} - Get player by Id.
func (ps *PlayerService) GetById(id int) playerStruct.Player {
	allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

	var newPlayer playerStruct.Player

	for _, player := range allPlayers {
		if player.PlayerId == id {
			newPlayer = player
		}
	}

	return newPlayer
}

// Request Service - PUT /players/{id} - Update player by Id.
func (ps *PlayerService) Update(id int, playerDTO playerStruct.Player) (playerStruct.Player, error) {
	if playerDTO.FirstName == "" || playerDTO.LastName == "" || playerDTO.PlayerId == 0 {
		return playerDTO, errors.New("complete the required fields")
	} else {
		allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

		var updatePlayer playerStruct.Player = playerDTO
		updatePlayer.PlayerId = id

		for i, player := range allPlayers {
			if player.PlayerId == id {
				player.FirstName = playerDTO.FirstName
				player.LastName = playerDTO.LastName
				player.TeamId = playerDTO.TeamId
				allPlayers = append(allPlayers[:i], player)
			}
		}

		byteValue, err := json.Marshal(allPlayers)
		if err != nil {
			logrus.New().Warn(err)
		}

		err = ioutil.WriteFile(playersJsonPath, byteValue, 0644)
		if err != nil {
			logrus.New().Warn(err)
		}

		return updatePlayer, nil
	}
}

// Request Service - DELETE /players/{id} - Delete player by Id.
func (ps *PlayerService) Delete(id int) playerStruct.Player {
	allPlayers := ps.helpers.UnmarshalPlayersJson(playersJsonPath)

	var deletedPlayer playerStruct.Player

	for i, player := range allPlayers {
		if player.PlayerId == id {
			allPlayers = append(allPlayers[:i], allPlayers[i+1:]...)
			deletedPlayer = player
		}
	}

	byteValue, err := json.Marshal(allPlayers)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(playersJsonPath, byteValue, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return deletedPlayer
}
