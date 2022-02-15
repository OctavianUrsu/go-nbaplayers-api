package main

import (
	"log"
	"os"

	api "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/handler"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/helpers"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/service"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/store"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Initializing config
	logrus.Infoln("Initializing config...")

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}

	port := os.Getenv("PORT")

	// Loading env variables
	logrus.Infoln("Loading env variables...")
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// Initialize connection to MongoDB
	db, err := store.NewMongoDB(store.Config{
		URI:        os.Getenv("DB_URI"),
		Name:       viper.GetString("db.name"),
		Password:   os.Getenv("DB_PASSWORD"),
		Collection: viper.GetString("db.collection"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}

	// Initialize dependency injection
	store := store.NewStore(db)
	helpers := new(helpers.Helpers)
	services := service.NewService(helpers, store)
	handlers := handler.NewHandler(services)

	// Starting the HTTP server
	logrus.Infof("Starting http server on port: %s\n", port)
	server := new(api.Server)

	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("An error occured while starting http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
