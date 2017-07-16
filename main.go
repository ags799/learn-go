package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	db, err := NewPostgresDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	id := uuid.NewV4()
	err = db.Add(Item{id, "walk dog"})
	if err != nil {
		fmt.Println(err)
		return
	}
	/*err = db.Remove(id)
	if err != nil {
		fmt.Println(err)
		return
	}*/
	NewGinServer(db).Run()
}

