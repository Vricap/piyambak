package database

import (
	"database/sql"
	"log"
)

func RunMigrations(DB *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user TEXT NOT NULL,
		message TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(schema)
	if err != nil {
		log.Fatalf("Error create table: %v", err)
	}
}
