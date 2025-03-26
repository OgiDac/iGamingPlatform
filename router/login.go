package router

import (
	"database/sql"
	"time"

	"github.com/gorilla/mux"
)

func NewLoginRouter(timeout time.Duration, db *sql.DB, r *mux.Router) {

	r.HandleFunc("/login", nil).Methods("POST")
}
