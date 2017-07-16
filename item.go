package main

import "github.com/satori/go.uuid"

type Item struct {
	Id uuid.UUID
	Description string `json:"description"`
}
