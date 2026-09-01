package database

import (
	"database/sql"
)

func RunMigrations(DB *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS rooms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    password TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    room_id INTEGER NOT NULL,
    user TEXT NOT NULL,
    message TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id)
        REFERENCES rooms(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
	);
	`

	_, err := DB.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
