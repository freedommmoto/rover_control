package api

import (
	"github.com/freedommmoto/rover_control/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) mapInfo(ctx *gin.Context) {
	mapSize, roverCommand, err := tool.ReadInstructionsFile(FilePart)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"map_size":   mapSize,
		"rover_step": len(roverCommand),
	})
}
