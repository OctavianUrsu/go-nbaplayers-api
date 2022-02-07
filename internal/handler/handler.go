package handler

import (
	mw "github.com/OctavianUrsu/go-nbaplayers-api/internal/middleware"
	"github.com/OctavianUrsu/go-nbaplayers-api/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
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
	logger := logrus.New()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(mw.NewStructuredLogger(logger))
	r.Use(middleware.Recoverer)

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
