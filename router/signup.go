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

func NewSignupRouter(env *config.Env, timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	pr := repository.NewPlayerRepository(db)
	sc := &controllers.SignupController{
		SignupUseCase: usecase.NewSignupUseCase(pr, timeout),
		Env:           env,
	}

	r.HandleFunc("/signup", sc.Signup).Methods("POST")
}
