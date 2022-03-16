package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Helpers struct{}

// Constructor for dependency injection
func NewHelper() *Helpers {
	return &Helpers{}
}

func (h *Helpers) UnmarshalPlayersJson(path string) []models.Player {
	// Read all players from the JSON file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	// Declare a variable to store all players
	var allPlayers []models.Player

	// Unmarshal the JSON file to allPlayers var
	json.Unmarshal(file, &allPlayers)

	return allPlayers
}

func (h *Helpers) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		logrus.Panic(err)
	}

	return string(bytes)
}

func (h *Helpers) VerifyPassword(hashedPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))

	if err != nil {
		return err == nil
	}

	return true
}

func (h *Helpers) GenerateToken(nickname string, password string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	var SECRET_KEY string = os.Getenv("SECRET_KEY")

	claims := &models.SignedClaims{
		Nickname: nickname,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}
