package repository

import (
	"context"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/jmoiron/sqlx"
)

type PlayerTournamentRepository interface {
	// MakeBet(ctx context.Context, request domain.PlayerTournamentRequest) (*domain.PlayerTournamentResponse, error)
	CreateOrUpdate(ctx context.Context, tx *sqlx.Tx, record domain.PlayerTournament) error
	GetRankingList(c context.Context, tournamentId int) ([]*domain.PlayerRankResponse, error)
}

type playerTournamentRepository struct {
	db *sqlx.DB
}

func (p *playerTournamentRepository) GetRankingList(c context.Context, tournamentId int) ([]*domain.PlayerRankResponse, error) {
	players := []*domain.PlayerRankResponse{}

	query := `
		SELECT
            players.id,
			players.name,
            RANK() OVER (ORDER BY player_tournaments.totalInvested DESC) AS playerRank,
			player_tournaments.totalInvested
        FROM player_tournaments JOIN players ON player_tournaments.player_id = players.id
        WHERE player_tournaments.tournament_id = ? 
	`

	err := p.db.Select(&players, query, tournamentId)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (p *playerTournamentRepository) CreateOrUpdate(ctx context.Context, tx *sqlx.Tx, record domain.PlayerTournament) error {
	query := `
			INSERT INTO player_tournaments (player_id, tournament_id, score, totalInvested)
			VALUES (:playerId, :tournamentId, :score, ABS(:score))
			ON DUPLICATE KEY UPDATE score = score + VALUES(score), totalInvested = totalInvested + ABS(:score)
			`
	_, err := tx.NamedExec(query, record)
	if err != nil {
		return err
	}
	return nil
}

func NewPlayerTournamentRepository(db *sqlx.DB) PlayerTournamentRepository {
	return &playerTournamentRepository{
		db: db,
	}
}
