package repository

import (
	"context"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type PlayerRepository interface {
	GetPlayerById(ctx context.Context, id int) (*domain.Player, error)
	GetPlayers(ctx context.Context) ([]*domain.Player, error)
	GetHighestEarners(ctx context.Context) ([]*domain.PlayerRankingResponse, error)
	UpdatePlayer(ctx context.Context, player *domain.Player) error
	DeletePlayer(ctx context.Context, id int) error
	CreatePlayer(ctx context.Context, user *domain.Player) (*domain.Player, error)
	GetPlayerByEmail(ctx context.Context, email string) (*domain.Player, error)
	UpdateAccountBalance(ctx context.Context, tx *sqlx.Tx, playerId int, amount float64) error
}

type playerRepository struct {
	db *sqlx.DB
}

func NewPlayerRepository(db *sqlx.DB) PlayerRepository {
	return &playerRepository{
		db: db,
	}
}

func (r *playerRepository) GetHighestEarners(ctx context.Context) ([]*domain.PlayerRankingResponse, error) {
	var players []*domain.PlayerRankingResponse
	err := r.db.Select(&players,
		`SELECT id, name, accountBalance, RANK() OVER (ORDER BY accountBalance DESC) AS playerRank
		  FROM players`)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (r *playerRepository) UpdateAccountBalance(ctx context.Context, tx *sqlx.Tx, playerId int, amount float64) error {
	var transaction *sqlx.Tx
	var err error

	if tx != nil {
		transaction = tx
	} else {
		transaction, err = r.db.Beginx()
		if err != nil {
			return err
		}
	}

	_, err = transaction.Exec("UPDATE players SET accountBalance = accountBalance + ? WHERE id = ?", amount, playerId)
	if err != nil {
		if tx == nil {
			transaction.Rollback()
		}
		return err
	}

	if tx == nil {
		return transaction.Commit()
	}
	return nil
}

func (r *playerRepository) GetPlayerByEmail(ctx context.Context, email string) (*domain.Player, error) {
	player := domain.Player{}
	err := r.db.Get(&player, `SELECT * FROM players WHERE email = ?`, email)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (r *playerRepository) GetPlayers(ctx context.Context) ([]*domain.Player, error) {
	var players []*domain.Player
	err := r.db.Select(&players, "SELECT * FROM players")
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (r *playerRepository) GetPlayerById(c context.Context, id int) (*domain.Player, error) {
	player := domain.Player{}
	err := r.db.Get(&player, `SELECT * FROM players WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (r *playerRepository) CreatePlayer(ctx context.Context, player *domain.Player) (*domain.Player, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	res, err := tx.NamedExec(`INSERT INTO players (name, email, password, accountBalance, role)
	 VALUES (:name, :email, :password, 0, 'user')`, player)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	player.Id = int(id)
	player.AccountBalance = 0
	player.Role = domain.PlayerRole("user")
	return player, nil
}

func (r *playerRepository) UpdatePlayer(ctx context.Context, player *domain.Player) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer tx.Commit()

	if player.Password != "" {
		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(player.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return err
		}

		player.Password = string(encryptedPassword)
	}

	fieldsQuery := ""

	if player.Email != "" {
		fieldsQuery += "email = :email,"
	}
	if player.Name != "" {
		fieldsQuery += "name = :name,"
	}
	if player.Password != "" {
		fieldsQuery += "password = :password,"
	}

	fieldsQuery = fieldsQuery[:len(fieldsQuery)-1]

	_, err = tx.NamedExec("UPDATE players SET "+fieldsQuery+" WHERE id = :id", player)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *playerRepository) DeletePlayer(ctx context.Context, id int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Commit()

	_, err = tx.Exec("DELETE FROM players WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
