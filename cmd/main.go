package main

import (
	"log"

	gonbaplayersapi "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/handler"
)

const PORT string = "8080"

func main() {
	log.Printf("Starting http server on port: %s", PORT)

	// object: handler instance
	handlers := new(handler.Handler)

	// object: server instance
	server := new(gonbaplayersapi.Server)

	// run server
	if err := server.Run(PORT, handlers.InitRoutes()); err != nil {
		log.Fatalf("An error occured while starting http server: %s", err.Error())
	}
}
