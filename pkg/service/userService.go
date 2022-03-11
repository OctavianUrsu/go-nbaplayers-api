package service

import (
	"errors"
	"os"
	"time"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/dgrijalva/jwt-go"
)

func (s *Service) SignUp(userSignupDTO structure.User) (*string, error) {
	// Create a hash password from the user password
	hashedPassword := s.helpers.HashPassword(userSignupDTO.Password)

	// Add the hashed password to the User structure and
	// add the date when the user was created
	userSignupDTO.Password = hashedPassword
	userSignupDTO.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// Generate an auth token
	token, err := s.helpers.GenerateToken(userSignupDTO.Nickname, userSignupDTO.Password)
	if err != nil {
		return nil, err
	}

	// Add the created user to the database
	err = s.store.Signup(&userSignupDTO)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *Service) SignIn(userSigninDTO structure.UserSignin) (*string, error) {
	// Find if user exists and get the hashed password
	userPassword, err := s.store.FindUserByNickname(userSigninDTO.Nickname)
	if err != nil {
		return nil, err
	}

	// Check if the hashed password coincides with the user password
	isPasswordCorrect, err := s.helpers.VerifyPassword(userPassword.Password, userSigninDTO.Password)
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	// If the password is correct generate an auth token
	// otherwise throw an error
	if isPasswordCorrect {
		return s.helpers.GenerateToken(userSigninDTO.Nickname, userPassword.Password)
	} else {
		return nil, errors.New("incorrect password")
	}
}

func (s *Service) VerifyToken(tokenString string) (*bool, error) {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	logrus.Fatal("Error loading .env file")
	// }

	var SECRET_KEY string = os.Getenv("SECRET_KEY")

	claims := &structure.SignedClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	userExists, err := s.store.FindUserByTokenClaims(claims)

	if err != nil {
		return userExists, err
	}

	return userExists, nil
}
