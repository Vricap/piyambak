package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() *sql.DB {
	var err error

	DB, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		log.Fatalf("Failed to load database: %v", err)
	}

	return DB
}
