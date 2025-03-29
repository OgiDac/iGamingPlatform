package domain

import "context"

type PlayerRole string

const (
	User  PlayerRole = "user"
	Admin PlayerRole = "admin"
)

type Player struct {
	Id             int        `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	Password       string     `json:"password" db:"password"`
	Email          string     `json:"email" db:"email"`
	AccountBalance float64    `json:"accountBalance" db:"accountBalance"`
	Role           PlayerRole `json:"role" db:"role"`
}

type PlayerResponse struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	AccountBalance float64 `json:"accountBalance"`
}
type PlayerRankingResponse struct {
	Id             int     `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	AccountBalance string  `json:"accountBalance" db:"accountBalance"`
	PlayerRank     float64 `json:"playerRank" db:"playerRank"`
}

type PlayerUseCase interface {
	GetPlayerById(c context.Context, id int) (*PlayerResponse, error)
	GetPlayers(c context.Context) ([]*PlayerResponse, error)
	UpdatePlayer(c context.Context, player *Player) error
	DeletePlayer(c context.Context, id int) error
	AddFunds(c context.Context, amount float64, playerId int) error
	GetHighestEarners(ctx context.Context) ([]*PlayerRankingResponse, error)
}
