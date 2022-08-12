package services

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestControlRover(t *testing.T) {
	//case basic N
	position := TwoDPosition{23, 0, 0}
	rover := RoverBasic{direction: "N", position2d: position}
	err := rover.ControlRover("F")
	require.NoError(t, err)

	//case not support command
	err = rover.ControlRover("B")
	require.Error(t, err)
	require.Equal(t, err.Error(), "command not support on basic rover")

	//case not support direction
	rover.direction = "NW"
	err = rover.ControlRover("F")
	require.Error(t, err)
	require.Equal(t, err.Error(), "direction not support")

	//case Advanced support command
	position = TwoDPosition{23, 0, 0}
	roverAdvanced := RoverAdvanced{direction: "N", position2d: position}
	err = roverAdvanced.ControlRover("B")
	require.NoError(t, err)

	//case Advanced support command
	roverAdvanced.direction = "NW"
	err = rover.ControlRover("F")
	require.Error(t, err)
	require.Equal(t, err.Error(), "direction not support")
}
