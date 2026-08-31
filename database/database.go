package database

import (
	"database/sql"
)

type dbContext struct {
	DB *sql.DB
}

var DbCtx *dbContext

func Connect() (*sql.DB, error) {
	var err error

	DB, err := sql.Open("sqlite3", "./database/chat.db")
	if err != nil {
		return nil, err
	}

	DbCtx = &dbContext{
		DB,
	}

	return DB, nil
}
