package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) mapInfo(ctx *gin.Context) {
	//tool.checkFileExists{}
	ctx.JSON(http.StatusOK, gin.H{
		"map_size":   24,
		"rover_step": 8,
	})
}
