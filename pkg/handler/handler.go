package handler

import (
	mw "github.com/OctavianUrsu/go-nbaplayers-api/pkg/middleware"
	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
}

// Constructor for dependency injection
func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://www.octavianursu.com/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Initialize API routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/players", func(r chi.Router) {
			r.Get("/", h.getAllPlayers) // Get all players
			r.Post("/", h.createPlayer) // Add new player

			// Use ID
			r.Get("/{id}", h.getPlayerById)   // Get player by id
			r.Put("/{id}", h.updatePlayer)    // Update player by id
			r.Delete("/{id}", h.deletePlayer) // Delete player by id

			// Search for name
			r.Get("/search", h.getPlayerByName) // Get player by name
		})

		r.Route("/teams", func(r chi.Router) {
			r.Get("/", h.getAllTeams) // Get all teams
			r.Post("/", h.createTeam) // Add new team

			// Use ID
			r.Get("/{id}", h.getTeamById)   // Get team by id
			r.Put("/{id}", h.updateTeam)    // Update team by id
			r.Delete("/{id}", h.deleteTeam) // Delete team by id
		})
	})

	return r
}
