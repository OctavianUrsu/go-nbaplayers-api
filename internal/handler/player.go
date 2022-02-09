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
	players := h.playerService.GetAll()

	// resp: Write all players as a JSON + write the http status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

// Request Handler - POST /players - Add new player.
func (h *Handler) createPlayer(w http.ResponseWriter, r *http.Request) {
	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.New().Warn(err)
	}

	// Create a Data Transfer Object from
	var playerDTO playerStruct.Player

	// Populate the DTO with our request
	json.Unmarshal(req, &playerDTO)

	// Create player
	newPlayer, err := h.playerService.Create(playerDTO)
	if err != nil {
		// resp: In case of error, write the error + the http status
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(err.Error()))
	} else {
		// resp: In case of success, write the created player + the http status
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newPlayer)
	}
}

// Request Handler - GET /players/{id} - Get player by Id.
func (h *Handler) getPlayerById(w http.ResponseWriter, r *http.Request) {
	// Get id from URL params and convert it to integer
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.New().Warn(err)
	}

	// Get player by id
	playerById := h.playerService.GetById(id)

	// resp: Write the player + the http status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playerById)
}

// Request Handler - PUT /players/{id} - Update player by Id.
func (h *Handler) updatePlayer(w http.ResponseWriter, r *http.Request) {
	// Get id from URL params and convert it to integer
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.New().Warn(err)
	}

	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.New().Warn(err)
	}

	// Create a Data Transfer Object from
	var playerDTO playerStruct.Player

	// Populate the DTO with our request
	json.Unmarshal(req, &playerDTO)

	// Update player
	updatedPlayer, err := h.playerService.Update(id, playerDTO)

	if err != nil {
		// resp: In case of error, write the error + the http status
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	} else {
		// resp: In case of success, write the updated player + the http status
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(updatedPlayer)
	}
}

// Request Handler - DELETE /players/{id} - Delete player by Id.
func (h *Handler) deletePlayer(w http.ResponseWriter, r *http.Request) {
	// Get id from URL params and convert it to integer
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.New().Warn(err)
	}

	// Delete player
	deletedPlayer := h.playerService.Delete(id)

	// resp: Write the deleted player + the http status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedPlayer)
}
