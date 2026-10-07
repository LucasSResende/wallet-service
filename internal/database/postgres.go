package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error) {

	db, err := sql.Open(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/wallet?sslmode=disable",
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
