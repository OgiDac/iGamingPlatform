package domain

import "context"

type PlayerTournament struct {
	PlayerId     int     `json:"playerId"      db:"playerId"`
	TournamentId int     `json:"tournamentId"  db:"tournamentId"`
	Score        float64 `json:"score" db:"score"`
}

type PlayerTournamentRequest struct {
	PlayerId     int     `json:"playerId"`
	TournamentId int     `json:"tournamentId"`
	Bet          float64 `json:"bet"`
}

type PlayerTournamentResponse struct {
	Result string `json:"result"`
}

type PlayerTournamentUseCase interface {
	MakeBet(c context.Context, playerTournamentRequest PlayerTournamentRequest) (*PlayerTournamentResponse, error)
}
