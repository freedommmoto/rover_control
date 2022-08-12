package services

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var mapSize = 24
var edge = mapSize - 1

func TestMoveNorth(t *testing.T) {
	//case normal
	position := TwoDPosition{edge, 0, 0}
	newPosition, err := moveNorth(position)
	require.NoError(t, err)
	require.NotEmpty(t, newPosition)
	require.Equal(t, newPosition.positionY, 1)

	//case have edge on y
	position = TwoDPosition{edge, 0, edge}
	newPosition, err = moveNorth(position)
	require.Error(t, err)
	require.Equal(t, err.Error(), "unable to move north is over the edge")
}
