package service

import (
	"errors"
	"os"
	"time"

	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func (s *Service) SignUp(userSignupDTO models.User) error {
	if s.store.NicknameExists(&userSignupDTO) {
		return errors.New("nickname already exists")
	}

	if s.store.EmailExists(&userSignupDTO) {
		return errors.New("email already exists")
	}

	// Create a hash password from the user password
	hashedPassword := s.helpers.HashPassword(userSignupDTO.Password)

	// Add the hashed password to the User models and
	// add the date when the user was created
	userSignupDTO.Password = hashedPassword
	userSignupDTO.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// Add the created user to the database
	err := s.store.Signup(&userSignupDTO)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SignIn(userSigninDTO models.UserSignin) (string, error) {
	// Find if user exists and get the hashed password
	hashedPassword, err := s.store.GetHashedPassword(userSigninDTO.Nickname)
	if err != nil {
		return "", err
	}

	// Check if the hashed password coincides with the user password
	isPasswordCorrect := s.helpers.VerifyPassword(hashedPassword, userSigninDTO.Password)
	if !isPasswordCorrect {
		return "", errors.New("wrong credentials")
	}

	// If the password is correct generate an auth token
	// otherwise throw an error
	if isPasswordCorrect {
		return s.helpers.GenerateToken(userSigninDTO.Nickname, hashedPassword)
	} else {
		return "", errors.New("wrong credentials")
	}
}

func (s *Service) VerifyToken(tokenString string) (bool, error) {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	var isTokenVerified bool = false
	var SECRET_KEY string = os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return isTokenVerified, err
	}

	if token.Valid {
		isTokenVerified = true
		return isTokenVerified, nil
	} else {
		return isTokenVerified, errors.New("invalid token")
	}
}
