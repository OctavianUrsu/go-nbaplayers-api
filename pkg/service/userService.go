package service

import (
	structure "github.com/OctavianUrsu/go-nbaplayers-api"
)

func (s *Service) SignUp(userSignupDTO structure.User) error {
	hashedPassword := s.helpers.HashPassword(userSignupDTO.Password)

	userSignupDTO.Password = hashedPassword

	if err := s.store.Signup(&userSignupDTO); err != nil {
		return err
	}

	return nil
}
