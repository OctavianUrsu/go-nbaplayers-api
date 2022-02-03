package handler

import (
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/service"
	"github.com/go-chi/chi/v5"
)

// class
type Handler struct {
	playerService *service.PlayerService
}

func NewHandler(ps *service.PlayerService) *Handler {
	return &Handler{ps}
}

// method
func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/players", func(r chi.Router) {
			r.Get("/", h.getAllPlayers)       // Get all players
			r.Post("/", h.createPlayer)       // Add new player
			r.Get("/{id}", h.getPlayerById)   // Get player by id
			r.Put("/{id}", h.updatePlayer)    // Update player by id
			r.Delete("/{id}", h.deletePlayer) // Delete player by id

		})
	})

	return r
}
