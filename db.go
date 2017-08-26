package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type db interface {
	List() ([]item, error)
	Add(item) error
	Remove(uuid.UUID) error
	Close() error
}

const tableName = "tasks"

type postgresDb struct {
	db     *sql.DB
	list   *sql.Stmt
	add    *sql.Stmt
	remove *sql.Stmt
}

func newPostgresDb() (db, error) {
	db, err := sql.Open("postgres", "host=db user=postgres sslmode=disable")
	//db, err := sql.Open("postgres", "postgres://postgres@db/postgres?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("Open error: %s", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Ping error: %s", err)
	}
	list, err := db.Prepare(fmt.Sprintf("select * from %s", tableName))
	if err != nil {
		return nil, fmt.Errorf("List statement preparation error: %s", err)
	}
	add, err := db.Prepare(fmt.Sprintf("insert into %s values ($1, $2)", tableName))
	if err != nil {
		return nil, fmt.Errorf("Add statement preparation error: %s", err)
	}
	remove, err := db.Prepare(fmt.Sprintf("delete from %s where id = $1", tableName))
	if err != nil {
		return nil, fmt.Errorf("Remove statement preparation error: %s", err)
	}
	return postgresDb{db, list, add, remove}, nil
}

func (db postgresDb) List() ([]item, error) {
	rows, err := db.list.Query()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error retrieving rows from Postgres database: %s", tableName), err)
	}
	defer rows.Close()
	var items []item
	for rows.Next() {
		var (
			id          uuid.UUID
			description string
		)
		err := rows.Scan(&id, &description)
		if err != nil {
			return nil, fmt.Errorf("Error scanning rows: %s", err)
		}
		items = append(items, item{ID: id, Description: description})
	}
	return items, nil
}

func (db postgresDb) Add(item item) error {
	_, err := db.add.Query(item.ID, item.Description)
	if err != nil {
		return fmt.Errorf("Error inserting item: %s", err)
	}
	return nil
}

func (db postgresDb) Remove(id uuid.UUID) error {
	_, err := db.remove.Query(id)
	if err != nil {
		return fmt.Errorf("Error deleting item: %s", err)
	}
	return nil
}

func (db postgresDb) Close() error {
	err := db.db.Close()
	if err != nil {
		return fmt.Errorf("Close error: %s", err)
	}
	return nil
}
