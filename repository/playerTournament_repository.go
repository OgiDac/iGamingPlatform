package repository

import (
	"context"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/jmoiron/sqlx"
)

type PlayerTournamentRepository interface {
	// MakeBet(ctx context.Context, request domain.PlayerTournamentRequest) (*domain.PlayerTournamentResponse, error)
	CreateOrUpdate(ctx context.Context, tx *sqlx.Tx, record domain.PlayerTournament) error
}

type playerTournamentRepository struct {
	db *sqlx.DB
}

func (p *playerTournamentRepository) CreateOrUpdate(ctx context.Context, tx *sqlx.Tx, record domain.PlayerTournament) error {
	query := `
			INSERT INTO player_tournaments (player_id, tournament_id, score)
			VALUES (:playerId, :tournamentId, :score)
			ON DUPLICATE KEY UPDATE score = score + VALUES(score)
			`
	_, err := tx.NamedExec(query, record)
	if err != nil {
		return err
	}
	return nil
}

// func (p *playerTournamentRepository) MakeBet(ctx context.Context, request domain.PlayerTournamentRequest) (*domain.PlayerTournamentResponse, error) {
// 	player := domain.Player{}
// 	tournament := domain.Tournament{}
// 	err := p.db.Get(&player, `SELECT * FROM players WHERE id = ?`, request.PlayerId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = p.db.Get(&tournament, `SELECT * FROM tournaments WHERE id = ?`, request.TournamentId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if request.Bet > player.AccountBalance {
// 		return nil, errors.New("account ballance not sufficient")
// 	}

// 	tx, err := p.db.Beginx()
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer tx.Commit()

// 	var exists bool
// 	err = p.db.Get(&exists, `SELECT EXISTS(SELECT 1 FROM player_tournaments WHERE player_id = ? AND tournament_id = ?)`, request.PlayerId, request.TournamentId)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	if r.Float64() < float64(tournament.ChanceToWin)/100 {
// 		// win
// 		if !exists {
// 			_, err := tx.NamedExec(`INSERT INTO player_tournaments (player_id, tournament_id, score) VALUES (:playerId, :tournamentId, :bet)`, player)
// 			if err != nil {
// 				tx.Rollback()
// 				return nil, err
// 			}
// 		} else {
// 			_, err := tx.NamedExec(`
// 				UPDATE player_tournaments
// 				SET score = score + :bet
// 				WHERE player_id = :playerId AND tournament_id = :tournamentId
// 			`, request)
// 			if err != nil {
// 				tx.Rollback()
// 				return nil, err
// 			}
// 		}

// 		return &domain.PlayerTournamentResponse{Result: "You won!!"}, nil
// 	} else {
// 		// lose
// 		if !exists {
// 			_, err := tx.NamedExec(`INSERT INTO player_tournaments (player_id, tournament_id, score) VALUES (:playerId, :tournamentId, -:bet)`, player)
// 			if err != nil {
// 				tx.Rollback()
// 				return nil, err
// 			}
// 		} else {
// 			_, err := tx.NamedExec(`
// 				UPDATE player_tournaments
// 				SET score = score - :bet
// 				WHERE player_id = :playerId AND tournament_id = :tournamentId
// 			`, request)
// 			if err != nil {
// 				tx.Rollback()
// 				return nil, err
// 			}
// 		}

// 		return &domain.PlayerTournamentResponse{Result: "You lost!!"}, nil
// 	}
// }

func NewPlayerTournamentRepository(db *sqlx.DB) PlayerTournamentRepository {
	return &playerTournamentRepository{
		db: db,
	}
}
