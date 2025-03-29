package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OgiDac/iGamingPlatform/config"
	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/utils"
)

type SignupController struct {
	SignupUseCase domain.SignupUseCase
	Env           *config.Env
}

func (sc *SignupController) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request domain.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken, err := sc.SignupUseCase.SignUp(ctx, request, sc.Env)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	response := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.JSON(w, http.StatusOK, response)
	return
}
