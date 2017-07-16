package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type Db interface {
	List() ([]Item, error)
	Add(Item) error
	Remove(uuid.UUID) error
	Close() error
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
	return PostgresDb{db}, nil
}

func (db PostgresDb) List() ([]Item, error) {
	rows, err := db.db.Query("select * from tasks")
	if err != nil {
		return nil, fmt.Errorf("Error retrieving tasks from Postgres database: %s", err)
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var (
			id uuid.UUID
			description string
		)
		err := rows.Scan(&id, &description)
		if err != nil {
			return nil, fmt.Errorf("Error scanning rows: %s", err)
		}
		items = append(items, Item{Id: id, Description: description})
	}
	return items, nil
}

func (db PostgresDb) Add(item Item) error {
	_, err := db.db.Query("insert into tasks values ($1, $2)", item.Id, item.Description)
	if err != nil {
		return fmt.Errorf("Error inserting item: %s", err)
	}
	return nil
}

func (db PostgresDb) Remove(id uuid.UUID) error {
	_, err := db.db.Query("delete from tasks where id = $1", id)
	if err != nil {
		return fmt.Errorf("Error deleting item: %s", err)
	}
	return nil
}

func (db PostgresDb) Close() error {
	err := db.db.Close()
	if err != nil {
		return fmt.Errorf("Close error: %s", err)
	}
	return nil
}
