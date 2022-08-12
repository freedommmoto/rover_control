package services

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var mapSize = 24
var edge = mapSize - 1

var iniX int
var iniY int

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
