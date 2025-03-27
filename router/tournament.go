package router

import (
	"time"

	"github.com/OgiDac/iGamingPlatform/api/controllers"
	"github.com/OgiDac/iGamingPlatform/repository"
	"github.com/OgiDac/iGamingPlatform/usecase"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewTournamentRouter(timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	tr := repository.NewTournamentRepository(db)
	tc := &controllers.TournamentController{
		TournamentUseCase: usecase.NewTournamentUseCase(tr, timeout),
	}

	group := r.PathPrefix("/tournament").Subrouter()

	group.HandleFunc("/{id}/distributePrizes", tc.ExecuteDistributePrizes).Methods("PUT")
}
