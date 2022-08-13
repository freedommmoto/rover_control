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
