package database

import (
	"database/sql"
)

func RunMigrations(DB *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user TEXT NOT NULL,
		message TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
