package tool

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var DemoRoute = []string{"R", "F", "L", "F", "L", "L", "F", "R"}

func TestReadInstructionsFileIntegration(t *testing.T) {
	//read and format from fail file
	scanner := loadFileScanner(t, "../instructions_file/mock_instructions_for_test_fail.txt")
	dataFromFile, err := ReadFile(scanner)
	require.NoError(t, err)
	require.NotEmpty(t, dataFromFile)

	_, _, err = FormatStingSliceToRoverCommand(dataFromFile)
	require.Equal(t, err.Error(), "command not long 1 character")
	require.Error(t, err)

	//read and format from ok file
	scannerOK := loadFileScanner(t, "../instructions_file/mock_instructions_for_test_ok.txt")
	dataFromFileOK, errOk := ReadFile(scannerOK)
	require.NoError(t, errOk)
	require.NotEmpty(t, dataFromFileOK)

	_, _, err = FormatStingSliceToRoverCommand(dataFromFileOK)
	require.NoError(t, err)
}

func TestReadInstructionsFile(t *testing.T) {
	mapSize, roverPart, err := ReadInstructionsFile("../instructions_file/instructions.txt")
	require.NoError(t, err)
	require.NotEmpty(t, mapSize)
	require.NotEmpty(t, roverPart)
	require.Equal(t, DemoRoute, roverPart)
}
