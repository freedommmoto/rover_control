package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

var server *Server
var FilePart = "../instructions_file/instructions.txt"

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/status", server.status)
	router.GET("/map-info", server.mapInfo)
	router.GET("/rover-status", server.roverStatus)
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET")

		if c.Request.Method == "OPTIONS" || c.Request.Method == "POST" {
			c.AbortWithStatus(404)
			return
		}

		c.Next()
	}
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
