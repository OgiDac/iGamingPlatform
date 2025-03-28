package domain

import (
	"context"

	"github.com/OgiDac/iGamingPlatform/config"
)

type LoginRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUseCase interface {
	Login(ctx context.Context, request LoginRequest, env *config.Env) (accessToken string, refreshToken string, err error)
}
