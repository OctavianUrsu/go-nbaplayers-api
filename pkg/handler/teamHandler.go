package handler

import (
	"encoding/json"
	"io"
	"net/http"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// Request Handler - GET /teams - Get all teams.
func (h *Handler) getAllTeams(w http.ResponseWriter, r *http.Request) {
	// Get all teams
	players, err := h.service.GetAllTeams()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// resp: Write all teams as a JSON + write the http status
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

// Request Handler - POST /teams - Add new player.
func (h *Handler) createTeam(w http.ResponseWriter, r *http.Request) {
	// Obtain the session token from the requests cookies
	c := getTokenFromCookie(w, r)

	isAuthorized, err := h.service.VerifyToken(c.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	if *isAuthorized {
		// Read the request
		req, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.New().Warn(err)
			return
		}

		// Create a Data Transfer Object from
		var teamDTO structure.Team

		// Populate the DTO with our request
		if err := json.Unmarshal(req, &teamDTO); err != nil {
			// resp: In case of error, write the error + the http status
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Validate the request input
		validateErr := validate.Struct(teamDTO)
		if validateErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(validateErr.Error()))
			return
		}

		// Create player
		if err := h.service.CreateTeam(teamDTO); err != nil {
			// resp: In case of error, write the error + the http status
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
			return
		}

		// resp: In case of success, write the created player + the http status
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, "team created")
	}
}

// Request Handler - GET /team/{id} - Get team by Id.
func (h *Handler) getTeamById(w http.ResponseWriter, r *http.Request) {
	// Get id from URL params and convert it to integer
	id := chi.URLParam(r, "id")

	// Get team by id
	teamById, err := h.service.GetTeamById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// resp: Write the player + the http status
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teamById)
}

// Request Handler - PUT /team/{id} - Update team by Id.
func (h *Handler) updateTeam(w http.ResponseWriter, r *http.Request) {
	// Obtain the session token from the requests cookies
	c := getTokenFromCookie(w, r)

	isAuthorized, err := h.service.VerifyToken(c.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	if *isAuthorized {
		// Get id from URL params and convert it to integer
		id := chi.URLParam(r, "id")

		// Read the request
		req, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.Warn(err)
			return
		}

		// Create a Data Transfer Object from
		var teamDTO structure.Team

		// Populate the DTO with our request
		if err := json.Unmarshal(req, &teamDTO); err != nil {
			// resp: In case of error, write the error + the http status
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Update player
		if err := h.service.UpdateTeam(id, teamDTO); err != nil {
			// resp: In case of error, write the error + the http status
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// resp: In case of success, write the updated player + the http status
		w.Header().Set("content-type", "application/text")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, "team updated")
	}
}

// Request Handler - DELETE /team/{id} - Delete team by Id.
func (h *Handler) deleteTeam(w http.ResponseWriter, r *http.Request) {
	// Obtain the session token from the requests cookies
	c := getTokenFromCookie(w, r)

	isAuthorized, err := h.service.VerifyToken(c.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	if *isAuthorized {
		// Get id from URL params
		id := chi.URLParam(r, "id")

		// Delete team
		if err := h.service.DeleteTeam(id); err != nil {
			// resp: In case of error, write the error + the http status
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// resp: In case of success, write the updated player + the http status
		w.Header().Set("content-type", "application/text")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, "team deleted")
	}
}
