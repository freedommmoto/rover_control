package tool

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFormatPositionRover(t *testing.T) {
	str := FormatPositionRover("N", 1, 1)
	require.NotEmpty(t, str)
	require.Equal(t, str, "N:1,1")
}
