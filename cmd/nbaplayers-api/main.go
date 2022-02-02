package main

import (
	"log"
	"net/http"

	"github.com/OctavianUrsu/go-nbaplayers-api/internal/players"
	"github.com/go-chi/chi/v5"
)

const PORT string = "8080"

func main() {
	log.Printf("Starting up on http://localhost:%s", PORT)

	r := chi.NewRouter()

	// Route and handle home page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NBA Players API"))
	})

	// Route and handle /players
	r.Mount("/players", players.PlayersResource{}.Routes())

	// Build HTTP Server
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
