package service

import (
	"errors"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
)

// Request Service - GET /teams - Get all teams.
func (s *Service) GetAllTeams() ([]*structure.Team, error) {
	allTeams, err := s.store.GetAllTeams()
	if err != nil {
		return nil, err
	}

	return allTeams, nil
}

// Request Service - POST /teams - Add new team.
func (s *Service) CreateTeam(teamDTO structure.Team) error {
	// Check if request has empty strings
	if teamDTO.Name == "" && teamDTO.Abbreviation == "" && teamDTO.Location == "" {
		return errors.New("complete the required fields")
	}

	if err := s.store.CreateTeam(&teamDTO); err != nil {
		return err
	}

	return nil
}

// Request Service - GET /teams/{id} - Get team by Id.
func (s *Service) GetTeamById(id string) (*structure.Team, error) {
	team, err := s.store.GetTeamById(id)
	if err != nil {
		return nil, err
	}

	return team, nil
}

// Request Service - PUT /team/{id} - Update team by Id.
func (s *Service) UpdateTeam(id string, teamDTO structure.Team) error {
	// Check if request has empty strings
	if teamDTO.Name == "" && teamDTO.Abbreviation == "" && teamDTO.Location == "" {
		return errors.New("complete the required fields")
	}

	if err := s.store.UpdateTeam(id, &teamDTO); err != nil {
		return err
	}

	return nil
}

// Request Service - DELETE /team/{id} - Delete team by Id.
func (s *Service) DeleteTeam(id string) error {
	if err := s.store.DeleteTeam(id); err != nil {
		return err
	}

	return nil
}
