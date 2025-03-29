package router

import (
	"time"

	"github.com/OgiDac/iGamingPlatform/api/controllers"
	"github.com/OgiDac/iGamingPlatform/repository"
	"github.com/OgiDac/iGamingPlatform/usecase"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewTournamentRouter(timeout time.Duration, db *sqlx.DB, admin *mux.Router, private *mux.Router) {
	tr := repository.NewTournamentRepository(db)
	tc := &controllers.TournamentController{
		TournamentUseCase: usecase.NewTournamentUseCase(tr, timeout),
	}

	adminGroup := admin.PathPrefix("/tournament").Subrouter()
	privateGroup := private.PathPrefix("/tournament").Subrouter()

	adminGroup.HandleFunc("/{id}/distributePrizes", tc.ExecuteDistributePrizes).Methods("PUT")
	privateGroup.HandleFunc("/", tc.GetAllTournaments).Methods("GET")
}
