package domain

import (
	"context"

	"github.com/OgiDac/iGamingPlatform/config"
)

type SignupRequest struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUseCase interface {
	SignUp(ctx context.Context, request SignupRequest, env *config.Env) (accessToken string, refreshToken string, err error)
}
