package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func main() {
	db, err := newPostgresDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	id := uuid.NewV4()
	err = db.Add(item{id, "walk dog"})
	if err != nil {
		fmt.Println(err)
		return
	}
	/*err = db.Remove(id)
	if err != nil {
		fmt.Println(err)
		return
	}*/
	newGinServer(db).Run()
}
