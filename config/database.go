package config

import (
	"database/sql"
	"fmt"
)

func NewDbConnection() *sql.DB {
	dsn := "root:1234@tcp(localhost:3306)/igaming"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connection to the database")
	}
	return db
}

func CloseDbConnection(db *sql.DB) {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		fmt.Println("Error closing the database")
	}

	fmt.Println("Database connection closed")
}
