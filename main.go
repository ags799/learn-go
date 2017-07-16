package main

import (
	"github.com/gin-gonic/gin"
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
	err = db.Remove(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	items, err := db.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(items)
	/*r := gin.Default()
	r.GET("/item", List)
	r.GET("/ping", ping)
	r.Run()*/
}

func list(c *gin.Context) {
	var items []Item
	//items = append(items, Item{Id: 0, Description: "some description"}, Item{Id: 1, Description: "another description"})
	c.JSON(200, items)
}

func ping(c *gin.Context) {
	c.String(200, "pong")
}
