package controllers

import (
	"context"
	"net/http"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/utils"
)

type PlayerController struct {
	PlayerUseCase domain.PlayerUseCase
}

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
