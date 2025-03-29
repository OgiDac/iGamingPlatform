package usecase

import (
	"context"
	"time"

	"github.com/OgiDac/iGamingPlatform/config"
	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/repository"
	"github.com/OgiDac/iGamingPlatform/utils"
	"golang.org/x/crypto/bcrypt"
)

type signupUseCase struct {
	playerRepository repository.PlayerRepository
	contextTimeout   time.Duration
}

func (s *signupUseCase) SignUp(ctx context.Context, request domain.SignupRequest, env *config.Env) (accessToken string, refreshToken string, err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", "", err
	}

	request.Password = string(encryptedPassword)

	player := &domain.Player{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	player, err = s.playerRepository.CreatePlayer(ctx, player)

	if err != nil {
		return "", "", err
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

func NewSignupUseCase(playerRepository repository.PlayerRepository, timeout time.Duration) domain.SignupUseCase {
	return &signupUseCase{
		playerRepository: playerRepository,
		contextTimeout:   timeout,
	}
}
