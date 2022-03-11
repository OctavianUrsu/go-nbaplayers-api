package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Helpers struct{}
type SignedClaims struct {
	Nickname string
	Password string
	jwt.StandardClaims
}

// Constructor for dependency injection
func NewHelper() *Helpers {
	return &Helpers{}
}

func (h *Helpers) UnmarshalPlayersJson(path string) []playerStruct.Player {
	// Read all players from the JSON file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	// Declare a variable to store all players
	var allPlayers []playerStruct.Player

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

func (h *Helpers) VerifyPassword(userPassword string, providedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	if err != nil {
		return false, err
	}

	return true, nil
}

func (h *Helpers) GenerateToken(nickname string, password string) (*string, error) {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	logrus.Fatal("Error loading .env file")
	// }

	var SECRET_KEY string = os.Getenv("SECRET_KEY")

	claims := &SignedClaims{
		Nickname: nickname,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, err
	}

	return &token, nil
}
