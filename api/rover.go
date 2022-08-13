package api

import (
	service "github.com/freedommmoto/rover_control/service"
	"github.com/freedommmoto/rover_control/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (server *Server) roverStatus(ctx *gin.Context) {

	step, exit := ctx.GetQuery("step")
	intStep, err := strconv.Atoi(step)
	if !exit || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Need 'step' parameter to load rover status.",
		})
		return
	}

	mapSize, roverCommand, err := tool.ReadInstructionsFile("../instructions_file/instructions.txt")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if intStep < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "step is lower rover first step. please try start as step 0.",
		})
		return
	}
	if intStep > len(roverCommand) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "step is over rover last step. please try lower step.",
		})
		return
	}

	roverStatus, errMoveRover := service.ControlRoverFromCurrentStep(intStep, mapSize, roverCommand)
	if errMoveRover != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": errMoveRover.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, roverStatus)
}
