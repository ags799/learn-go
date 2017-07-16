package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Db interface {
	List() ([]Item, error)
}

type PostgresDb struct {
	db *sql.DB
}

func NewPostgresDb() (Db, error) {
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
	return PostgresDb{db}, nil
}

func (db PostgresDb) List() ([]Item, error) {
	rows, err := db.db.Query("select * from tasks")
	if err != nil {
		return nil, fmt.Errorf("Error retrieving tasks from Postgres database: %s", err)
	}
	defer rows.Close()
	// TODO: parse the rows into []Item
	var items []Item
	return items, nil
}
