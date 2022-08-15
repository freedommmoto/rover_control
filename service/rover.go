package services

import (
	"errors"
)

var speedRover = 1

var mapTurnLeft90Degree = map[string]string{
	"N": "W",
	"W": "S",
	"S": "E",
	"E": "N",
}
var mapTurnRight90Degree = map[string]string{
	"N": "E",
	"E": "S",
	"S": "W",
	"W": "N",
}

var mapTurnLeft45Degree = map[string]string{
	"N":  "NW",
	"NW": "W",
	"W":  "SW",
	"SW": "S",
	"S":  "SE",
	"SE": "E",
	"E":  "NE",
	"NE": "N",
}

var mapTurnRight45Degree = map[string]string{
	"N":  "NE",
	"NE": "E",
	"E":  "SE",
	"SE": "S",
	"S":  "SW",
	"SW": "W",
	"W":  "NW",
	"NW": "N",
}
var DemoRoute = []string{"R", "F", "L", "F", "L", "L", "F", "R"}

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

func moveNorth(po TwoDPosition, speed int) (TwoDPosition, error) {
	if po.positionY+1 > po.edge {
		return po, errors.New("unable to move north is over the edge")
	}
	po.positionY = po.positionY + speed
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

func validationDirection45Degree(direction string) error {
	if _, ok := mapTurnLeft45Degree[direction]; !ok {
		return errors.New("direction not support")
	}
	return nil
}

func turnLeft90Degree(direction string) (string, error) {
	err := validationDirection90Degree(direction)
	if err != nil {
		return direction, err
	}
	return mapTurnLeft90Degree[direction], nil
}

func turnRight90Degree(direction string) (string, error) {
	err := validationDirection90Degree(direction)
	if err != nil {
		return direction, err
	}
	return mapTurnRight90Degree[direction], nil
}

func turnLeft45Degree(direction string) (string, error) {
	err := validationDirection45Degree(direction)
	if err != nil {
		return direction, err
	}
	return mapTurnLeft45Degree[direction], nil
}

func turnRight45Degree(direction string) (string, error) {
	err := validationDirection45Degree(direction)
	if err != nil {
		return direction, err
	}
	return mapTurnRight45Degree[direction], nil
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
