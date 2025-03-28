package usecase

import (
	"context"
	"errors"
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

func (pu *playerTournamentUseCase) GetRankingList(c context.Context, tournamentId int) ([]*domain.PlayerRankResponse, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	players, err := pu.playerTournamentRepository.GetRankingList(ctx, tournamentId)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (p *playerTournamentUseCase) MakeBet(c context.Context, playerTournamentRequest domain.PlayerTournamentRequest) (*domain.PlayerTournamentResponse, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	player, err := p.playerRepository.GetPlayerById(ctx, playerTournamentRequest.PlayerId)
	if err != nil {
		return nil, err
	}

	if player.AccountBalance < playerTournamentRequest.Bet {
		return nil, errors.New("you don't have enough money")
	}

	if playerTournamentRequest.Bet <= 0 {
		return nil, errors.New("invalid betting amount")
	}

	tournament, err := p.tournamentRepository.GetTournamentById(ctx, playerTournamentRequest.TournamentId)
	if err != nil {
		return nil, err
	}

	if time.Now().Before(tournament.StartDate) {
		return nil, errors.New("tournament has not started yet")
	}

	if time.Now().After(tournament.EndDate) {
		return nil, errors.New("tournament has finished")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var playerTorunamentRecord domain.PlayerTournament
	var playerTorunamentResponse domain.PlayerTournamentResponse
	tx, err := p.db.Beginx()
	defer tx.Commit()
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

func NewPlayerTournamentUseCase(
	playerTournamentRepository repository.PlayerTournamentRepository,
	playerRepository repository.PlayerRepository,
	tournamentRepository repository.TournamentRepository,
	contextTimeout time.Duration,
	db *sqlx.DB) domain.PlayerTournamentUseCase {
	return &playerTournamentUseCase{
		playerTournamentRepository: playerTournamentRepository,
		tournamentRepository:       tournamentRepository,
		playerRepository:           playerRepository,
		contextTimeout:             contextTimeout,
		db:                         db,
	}
}
