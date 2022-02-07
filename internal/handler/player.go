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
	w.WriteHeader(http.StatusOK)

	players := h.playerService.GetAll()
	json.NewEncoder(w).Encode(players)

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
		w.WriteHeader(http.StatusNotAcceptable)
		io.WriteString(w, "complete the required fields")
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(player)
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
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

	w.WriteHeader(http.StatusGone)
	json.NewEncoder(w).Encode(player)
}
