package main

import (
	"log"

	api "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/handler"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/helpers"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/service"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.Infof("Initializing config...")

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}

	var port string = viper.GetString("port")

	logrus.Infof("Starting http server on port: %s", port)

	storage := new(storage.Storage)
	helpers := new(helpers.Helpers)
	services := service.NewService(helpers, storage)
	handlers := handler.NewHandler(services)
	server := new(api.Server)

	// run server from server.go
	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("An error occured while starting http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
