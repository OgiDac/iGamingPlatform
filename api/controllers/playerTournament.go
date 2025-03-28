package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/utils"
	"github.com/gorilla/mux"
)

type PlayerTorunamentController struct {
	PlayerTournamentUseCase domain.PlayerTournamentUseCase
}

func (ptc *PlayerTorunamentController) MakeBet(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "player_id", r.Context().Value("player_id"))
	playerIdString := fmt.Sprintf("%v", ctx.Value("user_id"))

	playerId, err := strconv.Atoi(playerIdString)
	pathVars := mux.Vars(r)
	tournamentIdString := pathVars["tournamentId"]

	tournamentId, err := strconv.Atoi(tournamentIdString)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	var body struct {
		Bet float64 `json:"bet"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	playerTorunamentRequest := domain.PlayerTournamentRequest{
		PlayerId:     playerId,
		TournamentId: tournamentId,
		Bet:          body.Bet,
	}
	response, err := ptc.PlayerTournamentUseCase.MakeBet(ctx, playerTorunamentRequest)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, response)
	return
}

func (ptc *PlayerTorunamentController) GetRankingList(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "player_id", r.Context().Value("player_id"))

	pathVars := mux.Vars(r)
	tournamentIdString := pathVars["tournamentId"]

	tournamentId, err := strconv.Atoi(tournamentIdString)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	players, err := ptc.PlayerTournamentUseCase.GetRankingList(ctx, tournamentId)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, players)
	return
}
