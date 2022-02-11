package service

import (
	"errors"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/helpers"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/storage"
)

type PlayerService struct {
	helpers *helpers.Helpers
	storage *storage.Storage
}

// Constructor for dependency injection
func NewService(h *helpers.Helpers, r *storage.Storage) *PlayerService {
	return &PlayerService{h, r}
}

// Request Service - GET /players - Get all players.
func (ps *PlayerService) GetAll() ([]*playerStruct.Player, error) {
	allPlayers, err := ps.storage.GetAll()
	if err != nil {
		return nil, err
	}

	return allPlayers, nil
}

// Request Service - POST /players - Add new player.
func (ps *PlayerService) Create(playerDTO playerStruct.Player) error {
	// Check if request has empty strings
	if playerDTO.FirstName == "" && playerDTO.LastName == "" {
		return errors.New("complete the required fields")
	}

	if err := ps.storage.Create(&playerDTO); err != nil {
		return err
	}

	return nil
}

// Request Service - GET /players/{id} - Get player by Id.
func (ps *PlayerService) GetById(id int) (*playerStruct.Player, error) {
	player, err := ps.storage.GetById(id)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// Request Service - PUT /players/{id} - Update player by Id.
func (ps *PlayerService) Update(id int, playerDTO playerStruct.Player) error {
	// Check if request has empty strings
	if playerDTO.FirstName == "" && playerDTO.LastName == "" && playerDTO.PlayerId == 0 {
		return errors.New("complete the required fields")
	}

	if err := ps.storage.Update(id, &playerDTO); err != nil {
		return err
	}

	return nil

}

// Request Service - DELETE /players/{id} - Delete player by Id.
func (ps *PlayerService) Delete(id int) error {
	if err := ps.storage.Delete(id); err != nil {
		return err
	}

	return nil
}
