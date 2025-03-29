package router

import (
	"net/http"
	"time"

	"github.com/OgiDac/iGamingPlatform/api/middleware"
	"github.com/OgiDac/iGamingPlatform/config"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// @Summary      Health Check
// @Description  Just returns OK
// @Tags         test
// @Success      200  {string}  string  "ok"
// @Router       /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func Setup(env *config.Env, timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	public := r.PathPrefix("/public/api").Subrouter()
	private := r.PathPrefix("/private/api").Subrouter()
	admin := r.PathPrefix("/admin/api").Subrouter()

	private.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	admin.Use(middleware.AdminAuthMiddleware(env.AccessTokenSecret))

	NewPlayerRouter(timeout, db, public, private)
	NewTournamentRouter(timeout, db, admin, private)
	NewLoginRouter(env, timeout, db, public)
	NewSignupRouter(env, timeout, db, public)
	NewPlayerTournamentRouter(timeout, db, private)
}
