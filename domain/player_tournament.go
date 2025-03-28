package domain

import "context"

type PlayerTournament struct {
	PlayerId      int     `json:"playerId"      db:"playerId"`
	TournamentId  int     `json:"tournamentId"  db:"tournamentId"`
	Score         float64 `json:"score" db:"score"`
	TotalInvested float64 `json:"totalInvested" db:"totalInvested"`
}

type PlayerTournamentRequest struct {
	PlayerId     int     `json:"playerId"`
	TournamentId int     `json:"tournamentId"`
	Bet          float64 `json:"bet"`
}

type PlayerTournamentResponse struct {
	Result string `json:"result"`
}

type PlayerRankResponse struct {
	Id            int     `json:"id" db:"id"`
	Name          string  `json:"name" db:"name"`
	PlayerRank    int     `json:"playerRank" db:"playerRank"`
	TotalInvested float64 `json:"totalInvested" db:"totalInvested"`
}

type PlayerTournamentUseCase interface {
	MakeBet(c context.Context, playerTournamentRequest PlayerTournamentRequest) (*PlayerTournamentResponse, error)
	GetRankingList(c context.Context, tournamentId int) ([]*PlayerRankResponse, error)
}
