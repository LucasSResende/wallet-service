package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error) {

	connStr :=
		"host=localhost port=5432 user=postgres password=postgres dbname=wallet sslmode=disable"

	return sql.Open("postgres", connStr)
}