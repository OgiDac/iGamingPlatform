package router

import (
	"time"

	"github.com/OgiDac/iGamingPlatform/api/controllers"
	"github.com/OgiDac/iGamingPlatform/repository"
	"github.com/OgiDac/iGamingPlatform/usecase"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewPlayerRouter(timeout time.Duration, db *sqlx.DB, public *mux.Router, private *mux.Router) {
	pr := repository.NewPlayerRepository(db)
	pc := &controllers.PlayerController{
		PlayerUseCase: usecase.NewPlayerUseCase(pr, timeout),
	}

	publicGroup := public.PathPrefix("/player").Subrouter()
	privateGroup := private.PathPrefix("/player").Subrouter()

	publicGroup.HandleFunc("/players", pc.GetPlayers).Methods("GET")
	privateGroup.HandleFunc("/add-funds", pc.AddFunds).Methods("PUT")
	privateGroup.HandleFunc("/earners", pc.GetHighestEarners).Methods("GET")
}
