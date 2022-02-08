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
	w.WriteHeader(http.StatusOK)

	players := h.playerService.GetAll()
	json.NewEncoder(w).Encode(players)
}

// Request Handler - POST /players - Add new player.
func (h *Handler) createPlayer(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.New().Warn(err)
	}

	var playerDTO playerStruct.Player

	json.Unmarshal(req, &playerDTO)

	player, err := h.playerService.Create(playerDTO)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(player)
	}
}

// Request Handler - GET /players/{id} - Get player by Id.
func (h *Handler) getPlayerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.New().Warn(err)
	}

	player := h.playerService.GetById(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

// Request Handler - PUT /players/{id} - Update player by Id.
func (h *Handler) updatePlayer(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL Param
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.New().Warn(err)
	}

	req, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.New().Warn(err)
	}

	var playerDTO playerStruct.Player

	json.Unmarshal(req, &playerDTO)

	player, err := h.playerService.Update(id, playerDTO)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(player)
	}
}

// Request Handler - DELETE /players/{id} - Delete player by Id.
func (h *Handler) deletePlayer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		logrus.New().Warn(err)
	}

	player := h.playerService.Delete(id)

	w.WriteHeader(http.StatusGone)
	json.NewEncoder(w).Encode(player)
}
