package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	_, err := Postgres()
	if err != nil {
		fmt.Println(err)
		return
	}
	println("success")
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
