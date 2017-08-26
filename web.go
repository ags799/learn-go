package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type server interface {
	Run() error
}

type ginServer struct {
	engine *gin.Engine
	db     db
}

func newGinServer(db db) ginServer {
	engine := gin.Default()
	server := ginServer{engine, db}
	engine.GET("/item", server.List)
	engine.GET("/ping", server.Ping)
	return server
}

func (server ginServer) Run() error {
	return server.engine.Run()
}

func (server ginServer) List(c *gin.Context) {
	items, err := server.db.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, items)
}

func (server ginServer) Ping(c *gin.Context) {
	c.String(200, "pong")
}
