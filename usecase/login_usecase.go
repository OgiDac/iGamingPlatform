package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/OgiDac/iGamingPlatform/config"
	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/repository"
	"github.com/OgiDac/iGamingPlatform/utils"
	"golang.org/x/crypto/bcrypt"
)

type loginUseCase struct {
	playerRepository repository.PlayerRepository
	contextTimeout   time.Duration
}

func (l *loginUseCase) Login(ctx context.Context, request domain.LoginRequest, env *config.Env) (accessToken string, refreshToken string, err error) {

	player, err := l.playerRepository.GetPlayerByEmail(ctx, request.Email)
	if err != nil {
		return "", "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(request.Password)) != nil {
		return "", "", errors.New("invalid password")
	}
	accessToken, err = utils.CreateAccessToken(player, env.AccessTokenSecret, env.AccessTokenExpiryHour)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = utils.CreateRefreshToken(player, env.RefreshTokenSecret, env.RefreshTokenExpiryHour)
	if err != nil {
		return
	}

	return accessToken, refreshToken, nil
}

func NewLoginUseCase(playerRepository repository.PlayerRepository, timeout time.Duration) domain.LoginUseCase {
	return &loginUseCase{
		playerRepository: playerRepository,
		contextTimeout:   timeout,
	}
}
