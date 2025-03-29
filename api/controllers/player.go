package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/utils"
)

type PlayerController struct {
	PlayerUseCase domain.PlayerUseCase
}

// @Summary      Get Players
// @Description  Retrieve list of players
// @Tags         Player
// @Accept       json
// @Produce      json
// @Success      200 {array} domain.Player "List of players"
// @Failure      400 {object} domain.ErrorResponse "Bad request"
// @Router       /player/players [get]
func (pc *PlayerController) GetPlayers(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "player_id", r.Context().Value("player_id"))

	players, err := pc.PlayerUseCase.GetPlayers(ctx)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, players)
	return
}

// @Summary      Get Highest Earners
// @Description  Retrieve players with the highest earnings
// @Tags         Player
// @Accept       json
// @Produce      json
// @Success      200 {array} domain.Player "List of highest earners"
// @Failure      400 {object} domain.ErrorResponse "Bad request"
// @Router       /player/earners [get]
func (pc *PlayerController) GetHighestEarners(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "player_id", r.Context().Value("player_id"))

	players, err := pc.PlayerUseCase.GetHighestEarners(ctx)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, players)
	return
}

// @Summary      Add Funds
// @Description  Add funds to a player's account
// @Tags         Player
// @Accept       json
// @Produce      json
// @Success      204 {string} string "No Content"
// @Failure      400 {object} domain.ErrorResponse "Bad request"
// @Router       /player/add-funds [put]
func (pc *PlayerController) AddFunds(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "player_id", r.Context().Value("player_id"))

	var body struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	playerIdString := fmt.Sprintf("%v", ctx.Value("user_id"))

	playerId, err := strconv.Atoi(playerIdString)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = pc.PlayerUseCase.AddFunds(ctx, body.Amount, playerId)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
