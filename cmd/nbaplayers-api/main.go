package main

import (
	"io"
	"log"
	"net/http"

	"github.com/OctavianUrsu/go-nbaplayers-api/internal/players"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/teams"
)

const PORT string = ":8080"

func main() {
	// Route and handle home page
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/players", players.HandlePlayers)
	http.HandleFunc("/teams", teams.ReturnAllTeams)

	// Build HTTP Server
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home of NBA Players API")
}
