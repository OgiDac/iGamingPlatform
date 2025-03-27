package domain

import "context"

type Player struct {
	Id             int     `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	Password       string  `json:"password" db:"password"`
	Email          string  `json:"email" db:"email"`
	AccountBalance float64 `json:"accountBalance" db:"accountBalance"`
}

type PlayerResponse struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	AccountBalance float64 `json:"accountBalance"`
}

type PlayerUseCase interface {
	GetPlayerById(c context.Context, id int) (*PlayerResponse, error)
	GetPlayers(c context.Context) ([]*PlayerResponse, error)
	UpdatePlayer(c context.Context, player *Player) error
	DeletePlayer(c context.Context, id int) error
}
