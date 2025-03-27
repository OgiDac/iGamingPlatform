package repository

import (
	"context"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/jmoiron/sqlx"
)

type TournamentRepository interface {
	ExecuteDistributePrizes(c context.Context, id int) error
	GetTournamentById(c context.Context, id int) (*domain.Tournament, error)
}

type tournamentRepository struct {
	db *sqlx.DB
}

// GetTournamentById implements TournamentRepository.
func (t *tournamentRepository) GetTournamentById(c context.Context, id int) (*domain.Tournament, error) {
	tournament := domain.Tournament{}
	err := t.db.Get(&tournament, `SELECT * FROM tournaments WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return &tournament, nil
}

func (t *tournamentRepository) ExecuteDistributePrizes(c context.Context, id int) error {
	tx, err := t.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Commit()
	_, err = tx.Exec("CALL DistributePrizes(?)", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func NewTournamentRepository(db *sqlx.DB) TournamentRepository {
	return &tournamentRepository{
		db: db,
	}
}
