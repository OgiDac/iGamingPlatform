package router

import (
	"database/sql"
	"time"

	"github.com/gorilla/mux"
)

func NewUserRouter(timeout time.Duration, db *sql.DB, r *mux.Router) {

	group := r.PathPrefix("/user").Subrouter()

	group.HandleFunc("/list", nil).Methods("GET")
}
