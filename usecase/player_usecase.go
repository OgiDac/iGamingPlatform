package usecase

import (
	"context"
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

// DeletePlayer implements domain.PlayerUseCase.
func (pu *playerUseCase) DeletePlayer(c context.Context, id int) error {
	panic("unimplemented")
}

// GetPlayerById implements domain.PlayerUseCase.
func (pu *playerUseCase) GetPlayerById(c context.Context, id int) (*domain.PlayerResponse, error) {
	panic("unimplemented")
}

func (pu *playerUseCase) UpdatePlayer(c context.Context, player *domain.Player) error {
	panic("unimplemented")
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
