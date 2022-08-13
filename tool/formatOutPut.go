package tool

import (
	"errors"
	"strconv"
)

func FormatPositionRover(direction string, positionX int, positionY int) string {
	positionXStr := strconv.Itoa(positionX)
	positionYStr := strconv.Itoa(positionY)
	return direction + ":" + positionXStr + "," + positionYStr
}

func FormatStingSliceToRoverCommand(str []string) (mapSize int, route []string, err error) {
	allLine := make([]string, 0)

	for i, strLine := range str {
		if i == 0 {
			mapSize, err = strconv.Atoi(strLine)
			if mapSize < 1 {
				return 0, allLine, errors.New("first line is not map size")
			}
			if err != nil {
				return 0, allLine, err
			}
		} else {
			if len(strLine) == 1 {
				allLine = append(allLine, strLine)
			} else {
				return 0, nil, errors.New("command not long 1 character")
			}
		}
	}
	return mapSize, allLine, nil
}
