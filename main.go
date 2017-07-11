package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/item", list)
	r.GET("/ping", ping)
	r.Run()
}

type Item struct {
	Description string `json:"description"`
}

func list(c *gin.Context) {
	c.JSON(200, Item{Description: "some description"})
}

func ping(c *gin.Context) {
	c.String(200, "pong")
}
