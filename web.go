package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
}

type GinServer struct {
	engine *gin.Engine
	db     Db
}

func NewGinServer(db Db) GinServer {
	engine := gin.Default()
	server := GinServer{engine, db}
	engine.GET("/item", server.List)
	engine.GET("/ping", server.Ping)
	return server
}

func (server GinServer) Run() error {
	return server.engine.Run()
}

func (server GinServer) List(c *gin.Context) {
	items, err := server.db.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, items)
}

func (server GinServer) Ping(c *gin.Context) {
	c.String(200, "pong")
}
