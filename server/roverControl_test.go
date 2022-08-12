package services

import (
	"fmt"
	"github.com/freedommmoto/rover_control/tool"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestControlRoverCommandBasic(t *testing.T) {
	//case basic N forward
	position := TwoDPosition{edge, iniX, iniY}
	rover := RoverBasic{direction: "N", position2d: position}
	err := rover.ControlRover("F")
	require.NoError(t, err)

	//case not support command
	err = rover.ControlRover("B")
	require.Error(t, err)
	require.Equal(t, err.Error(), "command not support on basic rover")

	//case Basic not support direction
	rover.direction = "NW"
	err = rover.ControlRover("F")
	require.Error(t, err)
	require.Equal(t, err.Error(), "direction not support")

	//case direction not support
	rover.direction = "NW"
	err = rover.ControlRover("F")
	require.Error(t, err)
	require.Equal(t, err.Error(), "direction not support")
}

func TestControlRoverCommandAdvanced(t *testing.T) {
	//case Advanced support command
	position := TwoDPosition{edge, iniX, iniY}
	roverAdvanced := RoverAdvanced{direction: "N", position2d: position}
	err := roverAdvanced.ControlRover("B")
	require.NoError(t, err)

	//case not support command
	err = roverAdvanced.ControlRover("G")
	require.Error(t, err)
	require.Equal(t, err.Error(), "command not support on advanced rover")
}

func TestControlRoverMoveToMaxOfMap(t *testing.T) {
	position := TwoDPosition{edge, iniX, iniY}
	rover := RoverBasic{direction: "N", position2d: position}
	var err error
	for i := 0; i < edge; i++ {
		err = rover.ControlRover("F")
		require.NoError(t, err)
		err = rover.ControlRover("R")
		require.NoError(t, err)
		err = rover.ControlRover("F")
		require.NoError(t, err)
		err = rover.ControlRover("L")
		require.NoError(t, err)
	}
	require.Equal(t, rover.position2d, TwoDPosition{edge, edge, edge})

	//try move after on edge position should get error
	err = rover.ControlRover("F")
	require.Error(t, err)
}

func TestControlRoverWithDemoRoute(t *testing.T) {
	position := TwoDPosition{edge, iniX, iniY}
	rover := RoverBasic{direction: "N", position2d: position}

	route := [8]string{"R", "F", "L", "F", "L", "L", "F", "R"}
	routeOutputFormat := [9]string{"N:0,0", "E:0,0", "E:1,0", "N:1,0", "N:1,1", "W:1,1", "S:1,1", "S:1,0", "W:1,0"}
	var formatPosition string
	var err error

	for i, s := range route {
		err = rover.ControlRover(s)
		formatPosition = tool.FormatPositionRover(rover.direction, rover.position2d.positionX, rover.position2d.positionY)
		require.Equal(t, routeOutputFormat[i+1], formatPosition)
		fmt.Println(formatPosition)
		require.NoError(t, err)
	}
	positionAfterMove := rover.position2d
	require.Equal(t, positionAfterMove, TwoDPosition{edge, 1, 0})

}
