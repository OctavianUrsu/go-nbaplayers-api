package handler

import (
	"encoding/json"
	"io"
	"net/http"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// Request Handler - GET /players - Get all players.
func (h *Handler) getAllPlayers(w http.ResponseWriter, r *http.Request) {
	// Get all players
	players, err := h.service.GetAllPlayers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// resp: Write all players as a JSON + write the http status
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

// Request Handler - POST /players - Add new player.
func (h *Handler) createPlayer(w http.ResponseWriter, r *http.Request) {
	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.New().Warn(err)
		return
	}

	// Create a Data Transfer Object from
	var playerDTO structure.Player

	// Populate the DTO with our request
	json.Unmarshal(req, &playerDTO)

	// Validate the request input
	validateErr := validate.Struct(playerDTO)
	if validateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validateErr.Error()))
		return
	}

	// Create player
	if err := h.service.CreatePlayer(playerDTO); err != nil {
		// resp: In case of error, write the error + the http status
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(err.Error()))
		return
	}

	// resp: In case of success, write the created player + the http status
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "player created")
}

// Request Handler - GET /players/{id} - Get player by Id.
func (h *Handler) getPlayerById(w http.ResponseWriter, r *http.Request) {
	// Get id from URL params and convert it to integer
	id := chi.URLParam(r, "id")

	// Get player by id
	playerById, err := h.service.GetPlayerById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// resp: Write the player + the http status
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playerById)
}

// Request Handler - PUT /players/{id} - Update player by Id.
func (h *Handler) updatePlayer(w http.ResponseWriter, r *http.Request) {
	// Get id from URL params and convert it to integer
	id := chi.URLParam(r, "id")

	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Warn(err)
		return
	}

	// Create a Data Transfer Object from
	var playerDTO structure.Player

	// Populate the DTO with our request
	json.Unmarshal(req, &playerDTO)

	// Validate the request input
	validateErr := validate.Struct(playerDTO)
	if validateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validateErr.Error()))
		return
	}

	// Update player
	if err := h.service.UpdatePlayer(id, playerDTO); err != nil {
		// resp: In case of error, write the error + the http status
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// resp: In case of success, write the updated player + the http status
	w.Header().Set("content-type", "application/text")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "player updated")
}

// Request Handler - DELETE /players/{id} - Delete player by Id.
func (h *Handler) deletePlayer(w http.ResponseWriter, r *http.Request) {
	// Get id from URL params
	id := chi.URLParam(r, "id")

	// Delete player
	if err := h.service.DeletePlayer(id); err != nil {
		// resp: In case of error, write the error + the http status
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// resp: In case of success, write the updated player + the http status
	w.Header().Set("content-type", "application/text")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "player deleted")
}

// Request Handler - GET /players/?name={name} - Get player by name.
func (h *Handler) getPlayerByName(w http.ResponseWriter, r *http.Request) {
	// Get query params from url
	param := r.URL.Query()
	searchParam := param.Get("name")

	// Get player by name
	foundPlayers, err := h.service.GetPlayerByName(searchParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// resp: Write all players as a JSON + write the http status
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundPlayers)
}
