package tool

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckFileExists(t *testing.T) {
	//test load non file exits
	scanner, err := checkFileExists("mock_instructions_for_test_ok.txt")
	require.Error(t, err)
	require.Empty(t, scanner)

	loadFileScanner(t, "../instructions_file/mock_instructions_for_test_ok.txt")
}

func TestReadTextToSlice(t *testing.T) {
	nilArray := make([]string, 0)
	scanner := loadFileScanner(t, "../instructions_file/mock_instructions_for_test_ok.txt")
	dataFromFile, err := ReadFile(scanner)

	require.NoError(t, err)
	require.NotEqual(t, nilArray, dataFromFile, "should not be a nil array after read file")
	require.NotEmpty(t, dataFromFile)
}

func loadFileScanner(t *testing.T, path string) *FileScanner {
	scanner, err := checkFileExists(path)
	require.NoError(t, err)
	require.NotEmpty(t, scanner)
	return scanner
}
