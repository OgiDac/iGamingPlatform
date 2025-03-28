package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/utils"
	"github.com/gorilla/mux"
)

type TournamentController struct {
	TournamentUseCase domain.TournamentUseCase
}

func (tc *TournamentController) ExecuteDistributePrizes(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "player_id", r.Context().Value("player_id"))

	pathVars := mux.Vars(r)
	tournamentIdString := pathVars["id"]

	tournamentId, err := strconv.Atoi(tournamentIdString)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	err = tc.TournamentUseCase.ExecuteDistributePrizes(ctx, tournamentId)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
