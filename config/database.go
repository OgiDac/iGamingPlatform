package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewDbConnection(env *Env) *sqlx.DB {
	// dsn := "root:1234@tcp(db:3306)/igaming?parseTime=true"
	dsn := env.ConnString
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connection to the database")
	}
	return db
}

func CloseDbConnection(db *sqlx.DB) {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		fmt.Println("Error closing the database")
	}

	fmt.Println("Database connection closed")
}
