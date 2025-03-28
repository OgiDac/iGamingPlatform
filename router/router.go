package router

import (
	"time"

	"github.com/OgiDac/iGamingPlatform/api/middleware"
	"github.com/OgiDac/iGamingPlatform/config"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Setup(env *config.Env, timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	public := r.PathPrefix("/public/api").Subrouter()
	private := r.PathPrefix("/private/api").Subrouter()
	admin := r.PathPrefix("/admin/api").Subrouter()

	private.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	admin.Use(middleware.AdminAuthMiddleware(env.AccessTokenSecret))

	NewPlayerRouter(timeout, db, public)
	NewTournamentRouter(timeout, db, admin)
	NewLoginRouter(env, timeout, db, public)
	NewPlayerTournamentRouter(timeout, db, private)
}
