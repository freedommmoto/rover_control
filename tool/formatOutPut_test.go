package tool

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var SliceTestOkCase = []string{
	"10",
	"F",
	"R",
	"F",
	"L",
	"F",
}

var SliceTestFailCaseCommandNot1Carlong = []string{
	"10",
	"forward",
	"R",
}

var SliceTestFailMapSizeNotInt = []string{
	"R",
	"12",
	"F",
}

func TestFormatPositionRover(t *testing.T) {
	str := FormatPositionRover("N", 1, 1)
	require.NotEmpty(t, str)
	require.Equal(t, str, "N:1,1")
}

func TestFormatStingSliceToRoverCommand(t *testing.T) {
	mapSize, roverCommand, err := FormatStingSliceToRoverCommand(SliceTestOkCase)
	roverCommandTest := SliceTestOkCase[1:6]

	require.NotEmpty(t, mapSize)
	require.NotEmpty(t, roverCommand)
	require.NoError(t, err)
	require.Equal(t, roverCommand, roverCommandTest)

	//case fail
	_, _, err = FormatStingSliceToRoverCommand(SliceTestFailCaseCommandNot1Carlong)
	require.Error(t, err)

	_, _, err = FormatStingSliceToRoverCommand(SliceTestFailMapSizeNotInt)
	require.Error(t, err)
}
