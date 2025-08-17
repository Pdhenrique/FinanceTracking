package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(DB_USER string, DB_PASSWORD string, DB_HOST string, DB_PORT string, DB_NAME string) (*sql.DB, error) {

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME,
	)

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
