package router

import (
	"time"

	"github.com/OgiDac/iGamingPlatform/api/controllers"
	"github.com/OgiDac/iGamingPlatform/config"
	"github.com/OgiDac/iGamingPlatform/repository"
	"github.com/OgiDac/iGamingPlatform/usecase"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewLoginRouter(env *config.Env, timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	pr := repository.NewPlayerRepository(db)
	lc := &controllers.LoginController{
		LoginUseCase: usecase.NewLoginUseCase(pr, timeout),
		Env:          env,
	}
	r.HandleFunc("/login", lc.Login).Methods("POST")
}
