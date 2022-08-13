package services

import (
	"errors"
	"github.com/freedommmoto/rover_control/tool"
	"strings"
)

type RoverStatus struct {
	status      RoverBasic
	currentStep int    `json:"current_step"`
	statusSting string `json:"status_text_format"`
}

func ControlRoverFromCurrentStep(CurrentStep int, mapSize int, roverCommand []string) (rs RoverStatus, err error) {
	//init rover
	position := TwoDPosition{mapSize - 1, 0, 0}
	rover := RoverBasic{direction: "N", position2d: position}
	var formatPosition string

	//if first step no need to move rover
	if CurrentStep == 0 {
		rs.status = rover
		rs.statusSting = tool.FormatPositionRover(rover.direction, rover.position2d.positionX, rover.position2d.positionY)
		return rs, nil
	}

	//move rover after first step
	roverCommandThisStep := roverCommand[0:CurrentStep]

	for _, s := range roverCommandThisStep {
		err = rover.ControlRover(s)
		if err != nil {
			return rs, err
		}
		formatPosition = tool.FormatPositionRover(rover.direction, rover.position2d.positionX, rover.position2d.positionY)
	}

	rs.currentStep = CurrentStep
	rs.status = rover
	rs.statusSting = formatPosition
	return rs, nil
}

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
