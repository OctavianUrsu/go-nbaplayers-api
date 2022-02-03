package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	playerStruct "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/go-chi/chi/v5"
)

// Request Handler - GET /players - Get all players.
func (h *Handler) getAllPlayers(w http.ResponseWriter, r *http.Request) {
	players := h.playerService.GetAll()
	json.NewEncoder(w).Encode(players)

	w.WriteHeader(http.StatusOK)
}

// Request Handler - POST /players - Add new player.
func (h *Handler) createPlayer(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var playerDTO playerStruct.Player

	json.Unmarshal(req, &playerDTO)

	player, err := h.playerService.Create(playerDTO)
	if err != nil {
		io.WriteString(w, "complete the required fields")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	} else {
		json.NewEncoder(w).Encode(player)
		w.WriteHeader(http.StatusCreated)
	}
}

// Request Handler - GET /players/{id} - Get player by Id.
func (h *Handler) getPlayerById(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL Param
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
	}

	player := h.playerService.GetById(id)

	json.NewEncoder(w).Encode(player)
	w.WriteHeader(http.StatusOK)
}

// Request Handler - PUT /players/{id} - Update player by Id.
func (h *Handler) updatePlayer(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL Param
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	var playerDTO playerStruct.Player

	json.NewDecoder(r.Body).Decode(&playerDTO)
	if err != nil {
		fmt.Println(err)
	}

	player := h.playerService.Update(id, playerDTO)

	json.NewEncoder(w).Encode(player)
	w.WriteHeader(http.StatusOK)
}

// Request Handler - DELETE /players/{id} - Delete player by Id.
func (h *Handler) deletePlayer(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL Param
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	player := h.playerService.Delete(id)

	json.NewEncoder(w).Encode(player)
	w.WriteHeader(http.StatusOK)
}
