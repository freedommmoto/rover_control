package services

import (
	"errors"
	"strings"
)

func (rover *RoverBasic) ControlRover(NextCommand string) (err error) {
	direction := strings.ToUpper(rover.direction)

	if NextCommand == "F" {
		switch direction {
		case "N":
			rover.position2d, err = moveNorth(rover.position2d)
		case "W":
			rover.position2d, err = moveWest(rover.position2d)
		case "E":
			rover.position2d, err = moveEast(rover.position2d)
		case "S":
			rover.position2d, err = moveSouth(rover.position2d)
		default:
			err = errors.New("direction not support")
		}
	} else if NextCommand == "L" {
		rover.direction, err = turnLeft90Degree(direction)
	} else if NextCommand == "R" {
		rover.direction, err = turnRight90Degree(direction)
	} else {
		err = errors.New("command not support on basic rover")
	}

	if err != nil {
		return err
	}
	return nil
}

func (rover *RoverAdvanced) ControlRover(NextCommand string) (err error) {
	//direction := strings.ToUpper(rover.direction)

	if NextCommand == "F" {
		return nil
	} else if NextCommand == "B" {
		return nil
	} else if NextCommand == "L" {
		return nil
	} else if NextCommand == "R" {
		return nil
	} else if NextCommand == "HL" {
		return nil
	} else if NextCommand == "HR" {
		return nil
	} else {
		err = errors.New("command not support on advanced rover")
	}

	if err != nil {
		return err
	}
	return nil
}
