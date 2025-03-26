package router

import (
	"database/sql"
	"time"

	"github.com/gorilla/mux"
)

func Setup(timeout time.Duration, db *sql.DB, r *mux.Router) {
	public := r.PathPrefix("public/api").Subrouter()
	private := r.PathPrefix("private/api").Subrouter()

	NewLoginRouter(timeout, db, public)
	NewUserRouter(timeout, db, private)
}
