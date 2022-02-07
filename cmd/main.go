package main

import (
	api "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/handler"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/service"
	log "github.com/sirupsen/logrus"
)

const PORT string = "8080"

func main() {
	log.Info("Starting http server on port: ", PORT)

	// object: service instance
	services := new(service.PlayerService)

	// object: handler instance
	handlers := handler.NewHandler(services)

	// object: server instance
	server := new(api.Server)

	// run server
	if err := server.Run(PORT, handlers.InitRoutes()); err != nil {
		log.Fatal("An error occured while starting http server: ", err.Error())
	}
}
