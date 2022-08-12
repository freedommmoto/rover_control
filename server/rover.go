package services

import (
	"errors"
)

type RoverBasicControl interface {
	ControlRover(NextCommand string) error
}

type RoverBasic struct {
	direction  string
	position2d TwoDPosition
}

type RoverAdvanced struct {
	direction  string
	position2d TwoDPosition
}

type TwoDPosition struct {
	edge      int
	positionX int
	positionY int
}

func moveNorth(po TwoDPosition) (TwoDPosition, error) {
	if po.positionY+1 > po.edge {
		return po, errors.New("unable to move north is over the edge")
	}
	po.positionY++
	return po, nil
}

func moveSouth(po TwoDPosition) (TwoDPosition, error) {
	if po.positionY-1 < 0 {
		return po, errors.New("unable to move south is over the edge")
	}
	po.positionY--
	return po, nil
}

func moveEast(po TwoDPosition) (TwoDPosition, error) {
	if po.positionX+1 > po.edge {
		return po, errors.New("unable to move east is over the edge")
	}
	po.positionX++
	return po, nil
}

func moveWest(po TwoDPosition) (TwoDPosition, error) {
	if po.positionX-1 < 0 {
		return po, errors.New("unable to move west is over the edge")
	}
	po.positionX--
	return po, nil
}

func moveLeft90Degree(direction string) (string, error) {
	return direction, nil
}

func moveRight90Degree(direction string) (string, error) {
	return direction, nil
}

//todo work on it after done basic rover
func moveLeft45Degree(direction string) (string, error) {
	return direction, nil
}

func moveRight45Degree(direction string) (string, error) {
	return direction, nil
}

func moveNorthEast(po TwoDPosition) (TwoDPosition, error) {
	return po, nil
}
