package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

var server *Server

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/status", server.status)
	return router
}

func NewServer() *Server {
	router := setUpRouter()
	server = &Server{
		router: router,
	}
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
