package service

import (
	"errors"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignUp(userSignupDTO structure.User) error {
	if userSignupDTO.Nickname == "" && userSignupDTO.Email == "" && userSignupDTO.Password == "" {
		return errors.New("complete the required fields")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userSignupDTO.Password), 10)
	if err != nil {
		return err
	}

	userSignupDTO.Password = string(hashedPassword)

	if err := s.store.Signup(&userSignupDTO); err != nil {
		return err
	}

	return nil
}
