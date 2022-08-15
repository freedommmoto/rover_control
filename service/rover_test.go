package services

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var mapSize = 24
var edge = mapSize - 1

var iniX int
var iniY int

var iniDirection = "N"

func TestMoveNorth(t *testing.T) {
	//case normal
	position := TwoDPosition{edge, iniX, iniY}
	newPosition, err := moveNorth(position)
	require.NoError(t, err)
	require.NotEmpty(t, newPosition)
	require.Equal(t, newPosition.positionX, iniX)
	require.Equal(t, newPosition.positionY, iniY+1)

	//case have edge on y
	position = TwoDPosition{edge, iniX, edge}
	newPosition, err = moveNorth(position)
	require.Equal(t, newPosition.positionX, iniX)
	require.Equal(t, newPosition.positionY, edge)
	require.Error(t, err)
	require.Equal(t, err.Error(), "unable to move north is over the edge")
}

func TestMoveSouth(t *testing.T) {
	//case normal
	position := TwoDPosition{edge, iniX, edge}
	newPosition, err := moveSouth(position)
	require.NoError(t, err)
	require.NotEmpty(t, newPosition)
	require.Equal(t, newPosition.positionX, iniX)
	require.Equal(t, newPosition.positionY, edge-1)

	//case have edge on y
	position = TwoDPosition{edge, iniX, iniY}
	newPosition, err = moveSouth(position)
	require.Equal(t, newPosition.positionX, iniX)
	require.Equal(t, newPosition.positionY, iniY)
	require.Error(t, err)
	require.Equal(t, err.Error(), "unable to move south is over the edge")
}

func TestMoveEast(t *testing.T) {
	//case normal
	position := TwoDPosition{edge, iniX, iniY}
	newPosition, err := moveEast(position)
	require.NoError(t, err)
	require.NotEmpty(t, newPosition)
	require.Equal(t, newPosition.positionX, iniX+1)
	require.Equal(t, newPosition.positionY, iniY)

	//case have edge on x
	position = TwoDPosition{edge, edge, iniY}
	newPosition, err = moveEast(position)
	require.Equal(t, newPosition.positionX, edge)
	require.Equal(t, newPosition.positionY, iniY)
	require.Error(t, err)
	require.Equal(t, err.Error(), "unable to move east is over the edge")
}

func TestMoveWest(t *testing.T) {
	//case normal
	position := TwoDPosition{edge, edge, iniY}
	newPosition, err := moveWest(position)
	require.NoError(t, err)
	require.NotEmpty(t, newPosition)
	require.Equal(t, newPosition.positionX, edge-1)
	require.Equal(t, newPosition.positionY, iniY)

	//case have edge on x
	position = TwoDPosition{edge, iniX, iniY}
	newPosition, err = moveWest(position)
	require.Equal(t, newPosition.positionX, iniX)
	require.Equal(t, newPosition.positionY, iniY)
	require.Error(t, err)
	require.Equal(t, err.Error(), "unable to move west is over the edge")
}

func TestTurnLeft90Degree(t *testing.T) {
	//case un support
	_, errUnsupport := turnLeft90Degree("NW")
	require.Error(t, errUnsupport)

	//case normal
	direction, err := turnLeft90Degree(iniDirection)
	require.NoError(t, err)
	require.NotEmpty(t, direction)
	require.Equal(t, direction, "W")

	direction, err = turnLeft90Degree("W")
	require.NoError(t, err)
	require.NotEmpty(t, direction)
	require.Equal(t, direction, "S")

	direction, err = turnLeft90Degree("S")
	require.NoError(t, err)
	require.NotEmpty(t, direction)
	require.Equal(t, direction, "E")

	direction, err = turnLeft90Degree("E")
	require.NoError(t, err)
	require.NotEmpty(t, direction)
	require.Equal(t, direction, "N")

	//move 360 Degree back to init value
	for i := 0; i < 4; i++ {
		direction, _ = turnLeft90Degree(direction)
	}
	require.Equal(t, direction, iniDirection)
}

func TestTurnRight90Degree(t *testing.T) {
	//case un support
	_, errUnsupport := turnRight90Degree("NW")
	require.Error(t, errUnsupport)

	//case normal
	direction, err := turnRight90Degree(iniDirection)
	require.NoError(t, err)
	require.NotEmpty(t, direction)
	require.Equal(t, direction, "E")

	direction, err = turnRight90Degree("E")
	require.NoError(t, err)
	require.NotEmpty(t, direction)
	require.Equal(t, direction, "S")

	direction, err = turnRight90Degree("S")
	require.NoError(t, err)
	require.NotEmpty(t, direction)
	require.Equal(t, direction, "W")

	direction, err = turnRight90Degree("W")
	require.NoError(t, err)
	require.NotEmpty(t, direction)
	require.Equal(t, direction, "N")

	//move 360 Degree back to init value
	for i := 0; i < 4; i++ {
		direction, _ = turnRight90Degree(direction)
	}
	require.Equal(t, direction, iniDirection)
}

func TestTurnOver1cycle(t *testing.T) {
	direction := iniDirection
	for i := 0; i < 8; i++ {
		direction, _ = turnLeft90Degree(direction)
	}
	for j := 0; j < 8; j++ {
		direction, _ = turnRight90Degree(direction)
	}
	require.NotEmpty(t, direction)
	require.Equal(t, direction, "N")
}

//advanced rover

func TestTurnleft45Degree(t *testing.T) {
	//case un support on basic is support now on advanced
	_, errUnsupport := turnLeft45Degree("NE")
	require.NoError(t, errUnsupport)

	//case un support advanced
	_, errUnsupport = turnLeft45Degree("NN")
	require.Error(t, errUnsupport)

	direction := iniDirection
	//move 360 Degree back to init value
	for i := 0; i < 8; i++ {
		direction, _ = turnLeft45Degree(direction)
	}
	require.Equal(t, direction, iniDirection)
}

func TestTurnRight45Degree(t *testing.T) {
	//case un support on basic is support now on advanced
	_, errUnsupport := turnRight45Degree("NW")
	require.NoError(t, errUnsupport)

	//case un support advanced
	_, errUnsupport = turnRight45Degree("NN")
	require.Error(t, errUnsupport)

	direction := iniDirection
	//move 360 Degree back to init value
	for i := 0; i < 8; i++ {
		direction, _ = turnRight45Degree(direction)
	}
	require.Equal(t, direction, iniDirection)
}
