package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres sslmode=disable")
	if err != nil {
		print("Open: " + err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		print("Ping: " + err.Error())
		return
	}
	err = db.Close()
	if err != nil {
		print("Close: " + err.Error())
		return
	}
	print("success")
	/*r := gin.Default()
	r.GET("/item", list)
	r.GET("/ping", ping)
	r.Run()*/
}

type Item struct {
	Description string `json:"description"`
}

func list(c *gin.Context) {
	var items []Item
	items = append(items, Item{Description: "some description"}, Item{Description: "another description"})
	c.JSON(200, items)
}

func ping(c *gin.Context) {
	c.String(200, "pong")
}
