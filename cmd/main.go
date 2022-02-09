package main

import (
	"log"

	api "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/handler"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/service"
	"github.com/sirupsen/logrus"
)

const PORT string = "8080"

func main() {
	logrus.Infof("Starting http server on port: %s", PORT)

	// object: service instance
	services := new(service.PlayerService)

	// object: handler instance
	handlers := handler.NewHandler(services)

	// object: server instance
	server := new(api.Server)

	// run server from server.go
	if err := server.Run(PORT, handlers.InitRoutes()); err != nil {
		log.Fatal("An error occured while starting http server: ", err.Error())
	}
}
