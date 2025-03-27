package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/utils"
)

type LoginController struct {
	LoginUseCase domain.LoginUseCase
}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.JSON(w, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	accessToken, refreshToken, err := lc.LoginUseCase.Login(ctx, request)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	response := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.JSON(w, http.StatusOK, response)
	return
}
