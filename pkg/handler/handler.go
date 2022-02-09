package handler

import (
	mw "github.com/OctavianUrsu/go-nbaplayers-api/pkg/middleware"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	playerService *service.PlayerService
}

// Constructor for dependency injection
func NewHandler(ps *service.PlayerService) *Handler {
	return &Handler{ps}
}

func (h *Handler) InitRoutes() chi.Router {
	// Create a new logger
	logger := logrus.New()

	// Create a new router
	r := chi.NewRouter()

	// Use logger middleware
	r.Use(middleware.RequestID)
	r.Use(mw.NewStructuredLogger(logger))
	r.Use(middleware.Recoverer)

	// Initialize API routes
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
