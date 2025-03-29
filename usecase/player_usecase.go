package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/repository"
)

type playerUseCase struct {
	playerRepository repository.PlayerRepository
	contextTimeout   time.Duration
}

func NewPlayerUseCase(playerRepository repository.PlayerRepository, timeout time.Duration) domain.PlayerUseCase {
	return &playerUseCase{
		playerRepository: playerRepository,
		contextTimeout:   timeout,
	}
}

func (pu *playerUseCase) GetHighestEarners(ctx context.Context) ([]*domain.PlayerRankingResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	players, err := pu.playerRepository.GetHighestEarners(ctx)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (pu *playerUseCase) AddFunds(c context.Context, amount float64, playerId int) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	if amount <= 0 {
		return errors.New("invalid amount")
	}

	err := pu.playerRepository.UpdateAccountBalance(ctx, nil, playerId, amount)
	if err != nil {
		return err
	}

	return nil
}

// DeletePlayer implements domain.PlayerUseCase.
func (pu *playerUseCase) DeletePlayer(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	err := pu.playerRepository.DeletePlayer(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetPlayerById implements domain.PlayerUseCase.
func (pu *playerUseCase) GetPlayerById(c context.Context, id int) (*domain.PlayerResponse, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	player, err := pu.playerRepository.GetPlayerById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.PlayerResponse{
		Id:             player.Id,
		Name:           player.Name,
		Email:          player.Email,
		AccountBalance: player.AccountBalance,
	}, nil
}

func (pu *playerUseCase) UpdatePlayer(c context.Context, player *domain.Player) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	err := pu.playerRepository.UpdatePlayer(ctx, player)
	if err != nil {
		return err
	}

	return nil
}

func (pu *playerUseCase) GetPlayers(c context.Context) ([]*domain.PlayerResponse, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	var playersResponse []*domain.PlayerResponse
	players, err := pu.playerRepository.GetPlayers(ctx)
	if err != nil {
		return nil, err
	}

	for _, player := range players {
		playersResponse = append(playersResponse, &domain.PlayerResponse{
			Id:             player.Id,
			Email:          player.Email,
			Name:           player.Name,
			AccountBalance: player.AccountBalance,
		})
	}

	return playersResponse, nil
}
