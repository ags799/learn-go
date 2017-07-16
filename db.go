package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Postgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("Open error: %s", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Ping error: %s", err)
	}
	err = db.Close()
	if err != nil {
		return nil, fmt.Errorf("Close error: %s", err)
	}
	return db, nil
}
