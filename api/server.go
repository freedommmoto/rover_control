package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

var server *Server
var FilePart string

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/status", server.status)
	router.GET("/map-info", server.mapInfo)
	router.GET("/rover-status", server.roverStatus)
	return router
}

func NewServer() *Server {
	router := setUpRouter()
	server = &Server{
		router: router,
	}
	return server
}

func (server *Server) Start(address string, filePart string) error {
	FilePart = filePart
	return server.router.Run(address)
}
