package router

import (
	"time"

	"github.com/OgiDac/iGamingPlatform/api/controllers"
	"github.com/OgiDac/iGamingPlatform/repository"
	"github.com/OgiDac/iGamingPlatform/usecase"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewPlayerRouter(timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	pr := repository.NewPlayerRepository(db)
	pc := &controllers.PlayerController{
		PlayerUseCase: usecase.NewPlayerUseCase(pr, timeout),
	}

	group := r.PathPrefix("/player").Subrouter()

	group.HandleFunc("/players", pc.GetPlayers).Methods("GET")
}
