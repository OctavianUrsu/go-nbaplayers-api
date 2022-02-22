package service

import (
	"errors"
	"strings"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
)

// Request Service - GET /players - Get all players.
func (s *Service) GetAll() ([]*structure.Player, error) {
	allPlayers, err := s.store.GetAll()
	if err != nil {
		return nil, err
	}

	return allPlayers, nil
}

// Request Service - POST /players - Add new player.
func (s *Service) Create(playerDTO structure.Player) error {
	// Check if request has empty strings
	if playerDTO.FirstName == "" && playerDTO.LastName == "" {
		return errors.New("complete the required fields")
	}

	if err := s.store.Create(&playerDTO); err != nil {
		return err
	}

	return nil
}

// Request Service - GET /players/{id} - Get player by Id.
func (s *Service) GetById(id string) (*structure.Player, error) {
	player, err := s.store.GetById(id)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// Request Service - PUT /players/{id} - Update player by Id.
func (s *Service) Update(id string, playerDTO structure.Player) error {
	// Check if request has empty strings
	if playerDTO.FirstName == "" && playerDTO.LastName == "" {
		return errors.New("complete the required fields")
	}

	if err := s.store.Update(id, &playerDTO); err != nil {
		return err
	}

	return nil
}

// Request Service - DELETE /players/{id} - Delete player by Id.
func (s *Service) Delete(id string) error {
	if err := s.store.Delete(id); err != nil {
		return err
	}

	return nil
}

// Request Service - GET /players/?name={name} - Get player by name.
func (s *Service) GetByName(searchParam string) ([]*structure.Player, error) {
	searchParams := strings.Split(searchParam, " ")

	foundPlayers, err := s.store.GetByName(searchParams)
	if err != nil {
		return nil, err
	}

	return foundPlayers, nil
}
