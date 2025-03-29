package usecase

import (
	"context"
	"time"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/repository"
)

type tournamentUseCase struct {
	tournamentRepository repository.TournamentRepository
	contextTimeout       time.Duration
}

func (t *tournamentUseCase) GetAllTournaments(c context.Context) ([]*domain.Tournament, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	tournaments, err := t.tournamentRepository.GetTournaments(ctx)
	if err != nil {
		return nil, err

	}

	return tournaments, nil
}

func (t *tournamentUseCase) ExecuteDistributePrizes(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	err := t.tournamentRepository.ExecuteDistributePrizes(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func NewTournamentUseCase(tournamentRepository repository.TournamentRepository, contextTimeout time.Duration) domain.TournamentUseCase {
	return &tournamentUseCase{
		tournamentRepository: tournamentRepository,
		contextTimeout:       contextTimeout,
	}
}
