package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run()
}

func ping(c *gin.Context) {
	c.String(200, "pong")
}
