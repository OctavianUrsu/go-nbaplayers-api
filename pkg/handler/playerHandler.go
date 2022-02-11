package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// Request Handler - GET /players - Get all players.
func (h *Handler) getAllPlayers(w http.ResponseWriter, r *http.Request) {
	// Get all players
	players, err := h.playerService.GetAll()
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
	var playerDTO playerStruct.Player

	// Populate the DTO with our request
	json.Unmarshal(req, &playerDTO)

	// Create player
	if err := h.playerService.Create(playerDTO); err != nil {
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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Warn(err)
		return
	}

	// Get player by id
	playerById, err := h.playerService.GetById(id)
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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Warn(err)
		return
	}

	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Warn(err)
		return
	}

	// Create a Data Transfer Object from
	var playerDTO playerStruct.Player

	// Populate the DTO with our request
	json.Unmarshal(req, &playerDTO)

	// Update player
	if err := h.playerService.Update(id, playerDTO); err != nil {
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
	// Get id from URL params and convert it to integer
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.New().Warn(err)
	}

	// Delete player
	if err := h.playerService.Delete(id); err != nil {
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
