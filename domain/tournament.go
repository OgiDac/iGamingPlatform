package domain

import (
	"context"
	"time"
)

type Tournament struct {
	Id          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	PrizePool   float64   `json:"prizePool" db:"prizePool"`
	StartDate   time.Time `json:"startDate" db:"startDate"`
	EndDate     time.Time `json:"endDate" db:"endDate"`
	ChanceToWin int       `json:"chanceToWin" db:"chanceToWin"`
}

type TournamentResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	PrizePool   float64   `json:"prizePool"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	ChanceToWin int       `json:"chanceToWin" db:"chanceToWin"`
}

type TournamentUseCase interface {
	ExecuteDistributePrizes(c context.Context, id int) error
}
