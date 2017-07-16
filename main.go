package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	db, err := NewPostgresDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	items := db.List()
	fmt.Println(items)
	/*r := gin.Default()
	r.GET("/item", List)
	r.GET("/ping", ping)
	r.Run()*/
}

func list(c *gin.Context) {
	var items []Item
	items = append(items, Item{Description: "some description"}, Item{Description: "another description"})
	c.JSON(200, items)
}

func ping(c *gin.Context) {
	c.String(200, "pong")
}
