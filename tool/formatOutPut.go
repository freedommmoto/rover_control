package tool

import (
	"strconv"
)

func FormatPositionRover(direction string, positionX int, positionY int) string {
	positionXStr := strconv.Itoa(positionX)
	positionYStr := strconv.Itoa(positionY)
	return direction + ":" + positionXStr + "," + positionYStr
}
