package services

import (
	"errors"
)

var speedRover = 1

var mapTurnLeft = map[string]string{
	"N": "W",
	"W": "S",
	"S": "E",
	"E": "N",
}
var mapTurnRight = map[string]string{
	"N": "E",
	"E": "S",
	"S": "W",
	"W": "N",
}

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
	if po.positionY+speedRover > po.edge {
		return po, errors.New("unable to move north is over the edge")
	}
	po.positionY += speedRover
	return po, nil
}

func moveSouth(po TwoDPosition) (TwoDPosition, error) {
	if po.positionY-speedRover < 0 {
		return po, errors.New("unable to move south is over the edge")
	}
	po.positionY -= speedRover
	return po, nil
}

func moveEast(po TwoDPosition) (TwoDPosition, error) {
	if po.positionX+speedRover > po.edge {
		return po, errors.New("unable to move east is over the edge")
	}
	po.positionX += speedRover
	return po, nil
}

func moveWest(po TwoDPosition) (TwoDPosition, error) {
	if po.positionX-speedRover < 0 {
		return po, errors.New("unable to move west is over the edge")
	}
	po.positionX -= speedRover
	return po, nil
}

func validationDirection90Degree(direction string) error {
	switch direction {
	case "N":
		return nil
	case "W":
		return nil
	case "E":
		return nil
	case "S":
		return nil
	default:
		return errors.New("direction not support")
	}
}

func turnLeft90Degree(direction string) (string, error) {
	err := validationDirection90Degree(direction)
	if err != nil {
		return direction, err
	}
	return mapTurnLeft[direction], nil
}

func turnRight90Degree(direction string) (string, error) {
	err := validationDirection90Degree(direction)
	if err != nil {
		return direction, err
	}
	return mapTurnRight[direction], nil
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
