package usecase

import (
	"context"
	"math/rand"
	"time"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/repository"
	"github.com/jmoiron/sqlx"
)

type playerTournamentUseCase struct {
	playerTournamentRepository repository.PlayerTournamentRepository
	playerRepository           repository.PlayerRepository
	tournamentRepository       repository.TournamentRepository
	contextTimeout             time.Duration
	db                         *sqlx.DB
}

func (p *playerTournamentUseCase) MakeBet(c context.Context, playerTournamentRequest domain.PlayerTournamentRequest) (*domain.PlayerTournamentResponse, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	_, err := p.playerRepository.GetPlayerById(ctx, playerTournamentRequest.PlayerId)
	if err != nil {
		return nil, err
	}

	tournament, err := p.tournamentRepository.GetTournamentById(ctx, playerTournamentRequest.TournamentId)
	if err != nil {
		return nil, err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var playerTorunamentRecord domain.PlayerTournament
	var playerTorunamentResponse domain.PlayerTournamentResponse
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, err
	}
	if r.Float64() < float64(tournament.ChanceToWin)/100 {
		// win
		playerTorunamentRecord = domain.PlayerTournament{
			PlayerId:     playerTournamentRequest.PlayerId,
			TournamentId: playerTournamentRequest.TournamentId,
			Score:        playerTournamentRequest.Bet,
		}
		err = p.playerRepository.UpdateAccountBalance(ctx, tx, playerTournamentRequest.PlayerId, playerTournamentRequest.Bet)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		playerTorunamentResponse = domain.PlayerTournamentResponse{
			Result: "You won!!",
		}
	} else {
		// lose
		playerTorunamentRecord = domain.PlayerTournament{
			PlayerId:     playerTournamentRequest.PlayerId,
			TournamentId: playerTournamentRequest.TournamentId,
			Score:        -playerTournamentRequest.Bet,
		}
		err = p.playerRepository.UpdateAccountBalance(ctx, tx, playerTournamentRequest.PlayerId, -playerTournamentRequest.Bet)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		playerTorunamentResponse = domain.PlayerTournamentResponse{
			Result: "You lost!!",
		}
	}
	err = p.playerTournamentRepository.CreateOrUpdate(ctx, tx, playerTorunamentRecord)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &playerTorunamentResponse, nil
}

func NewPlayerTournamentUseCase(playerTournamentRepository repository.PlayerTournamentRepository, contextTimeout time.Duration) domain.PlayerTournamentUseCase {
	return &playerTournamentUseCase{
		playerTournamentRepository: playerTournamentRepository,
		contextTimeout:             contextTimeout,
	}
}
