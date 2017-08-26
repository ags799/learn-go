package main

import "github.com/satori/go.uuid"

type item struct {
	ID          uuid.UUID
	Description string `json:"description"`
}
