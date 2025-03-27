package router

import (
	"time"

	"github.com/OgiDac/iGamingPlatform/api/controllers"
	"github.com/OgiDac/iGamingPlatform/repository"
	"github.com/OgiDac/iGamingPlatform/usecase"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewPlayerTournamentRouter(timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	ptr := repository.NewPlayerTournamentRepository(db)
	pr := repository.NewPlayerRepository(db)
	tr := repository.NewTournamentRepository(db)
	ptc := &controllers.PlayerTorunamentController{
		PlayerTournamentUseCase: usecase.NewPlayerTournamentUseCase(ptr, pr, tr, timeout, db),
	}

	group := r.PathPrefix("/player-tournament").Subrouter()
	group.HandleFunc("/{tournamentId}/bet", ptc.MakeBet).Methods("PUT")
}
