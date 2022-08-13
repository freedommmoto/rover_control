package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"server_status": "ok",
	})
}
