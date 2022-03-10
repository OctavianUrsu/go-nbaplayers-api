package service

import (
	"errors"
	"strings"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
)

// Request Service - GET /players - Get all players.
func (s *Service) GetAllPlayers() ([]*structure.Player, error) {
	allPlayers, err := s.store.GetAllPlayers()
	if err != nil {
		return nil, err
	}

	return allPlayers, nil
}

// Request Service - POST /players - Add new player.
func (s *Service) CreatePlayer(playerDTO structure.Player) error {
	// Check if request has empty strings
	if playerDTO.FirstName == "" && playerDTO.LastName == "" {
		return errors.New("complete the required fields")
	}

	if err := s.store.CreatePlayer(&playerDTO); err != nil {
		return err
	}

	return nil
}

// Request Service - GET /players/{id} - Get player by Id.
func (s *Service) GetPlayerById(id string) (*structure.Player, error) {
	player, err := s.store.GetPlayerById(id)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// Request Service - PUT /players/{id} - Update player by Id.
func (s *Service) UpdatePlayer(id string, playerDTO structure.Player) error {
	// Check if request has empty strings
	if playerDTO.FirstName == "" && playerDTO.LastName == "" {
		return errors.New("complete the required fields")
	}

	if err := s.store.UpdatePlayer(id, &playerDTO); err != nil {
		return err
	}

	return nil
}

// Request Service - DELETE /players/{id} - Delete player by Id.
func (s *Service) DeletePlayer(id string) error {
	if err := s.store.DeletePlayer(id); err != nil {
		return err
	}

	return nil
}

// Request Service - GET /players/?name={name} - Get player by name.
func (s *Service) GetPlayerByName(searchParam string) ([]*structure.Player, error) {
	searchParams := strings.Split(searchParam, " ")

	foundPlayers, err := s.store.GetPlayerByName(searchParams)
	if err != nil {
		return nil, err
	}

	return foundPlayers, nil
}
