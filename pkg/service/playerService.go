package service

import (
	"errors"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/helpers"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/store"
)

type PlayerService struct {
	helpers *helpers.Helpers
	store   *store.Store
}

// Constructor for dependency injection
func NewService(h *helpers.Helpers, r *store.Store) *PlayerService {
	return &PlayerService{h, r}
}

// Request Service - GET /players - Get all players.
func (ps *PlayerService) GetAll() ([]*playerStruct.Player, error) {
	allPlayers, err := ps.store.GetAll()
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

	if err := ps.store.Create(&playerDTO); err != nil {
		return err
	}

	return nil
}

// Request Service - GET /players/{id} - Get player by Id.
func (ps *PlayerService) GetById(id string) (*playerStruct.Player, error) {
	player, err := ps.store.GetById(id)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// Request Service - PUT /players/{id} - Update player by Id.
func (ps *PlayerService) Update(id string, playerDTO playerStruct.Player) error {
	// Check if request has empty strings
	if playerDTO.FirstName == "" && playerDTO.LastName == "" {
		return errors.New("complete the required fields")
	}

	if err := ps.store.Update(id, &playerDTO); err != nil {
		return err
	}

	return nil
}

// Request Service - DELETE /players/{id} - Delete player by Id.
func (ps *PlayerService) Delete(id string) error {
	if err := ps.store.Delete(id); err != nil {
		return err
	}

	return nil
}

// Request Service - GET /players//?name={name} - Get player by name.
func (ps *PlayerService) GetByName(nameParam string) ([]*playerStruct.Player, error) {
	foundPlayers, err := ps.store.GetByName(nameParam)
	if err != nil {
		return nil, err
	}

	return foundPlayers, nil
}
