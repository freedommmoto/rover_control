package tool

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func checkFileExists(path string) (*FileScanner, error) {
	inFile, err := os.Open(path)
	if err != nil {
		return nil, errors.New(err.Error() + `: ` + path)
	}

	scanner := bufio.NewScanner(inFile)
	fileScanner := &FileScanner{inFile, scanner}
	return fileScanner, nil
}

func ReadFile(scanner *FileScanner) ([]string, error) {
	defer scanner.Close()
	allLine := make([]string, 0)
	var line string

	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		allLine = append(allLine, line)
	}
	return allLine, nil
}

func ReadInstructionsFile(part string) (mapSize int, route []string, err error) {
	scanner, err := checkFileExists(part)
	var dataFromFile []string
	if err != nil {
		return 0, nil, err
	}
	dataFromFile, err = ReadFile(scanner)
	if err != nil {
		return 0, nil, err
	}
	mapSize, route, err = FormatStingSliceToRoverCommand(dataFromFile)
	if err != nil {
		return 0, nil, err
	}

	return mapSize, route, nil
}
