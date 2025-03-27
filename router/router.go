package router

import (
	"time"

	"github.com/OgiDac/iGamingPlatform/api/middleware"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Setup(timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	public := r.PathPrefix("/public/api").Subrouter()
	private := r.PathPrefix("/private/api").Subrouter()

	private.Use(middleware.JwtAuthMiddleware("temp-secret"))

	NewPlayerRouter(timeout, db, public)
	NewTournamentRouter(timeout, db, private)
	NewLoginRouter(timeout, db, public)
	NewPlayerTournamentRouter(timeout, db, private)
}
