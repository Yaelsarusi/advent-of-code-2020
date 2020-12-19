package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type action int

const (
	east action = iota
	west
	north
	south
	left
	right
	forward
)

type instruction struct {
	action action
	value  int
}

func findElementIndex(directions []action, direction action) int {
	for i, val := range directions {
		if direction == val {
			return i
		}
	}
	return -1
}

func part1(instructions []instruction) {
	northSouthAxis := 0
	eastWestAxis := 0
	currentDirection := east
	directionOrder := []action{east, south, west, north}
	for _, instruction := range instructions {
		switch instruction.action {
		case east:
			eastWestAxis += instruction.value
		case west:
			eastWestAxis -= instruction.value
		case north:
			northSouthAxis += instruction.value
		case south:
			northSouthAxis -= instruction.value
		case left:
			currentPosition := findElementIndex(directionOrder, currentDirection)
			currentDirection = directionOrder[((64 + (currentPosition) - (instruction.value / 90)) % 4)]
		case right:
			currentPosition := findElementIndex(directionOrder, currentDirection)
			currentDirection = directionOrder[((64 + (currentPosition) + (instruction.value / 90)) % 4)]
		case forward:
			switch currentDirection {
			case east:
				eastWestAxis += instruction.value
			case west:
				eastWestAxis -= instruction.value
			case north:
				northSouthAxis += instruction.value
			case south:
				northSouthAxis -= instruction.value
			}
		}
	}
	fmt.Println(math.Abs(float64(northSouthAxis)) + math.Abs(float64(eastWestAxis)))
}

type wayPoint struct {
	northSouthAxis int
	eastWestAxis   int
}

func part2(instructions []instruction) {
	northSouthAxis := 0
	eastWestAxis := 0
	currentWayPoint := wayPoint{1, 10}
	for _, instruction := range instructions {
		switch instruction.action {
		case east:
			currentWayPoint.eastWestAxis += instruction.value
		case west:
			currentWayPoint.eastWestAxis -= instruction.value
		case north:
			currentWayPoint.northSouthAxis += instruction.value
		case south:
			currentWayPoint.northSouthAxis -= instruction.value
		case left:
			if instruction.value == 90 {
				newNorthSouthAxis := currentWayPoint.eastWestAxis
				newEastWestAxis := currentWayPoint.northSouthAxis * -1
				currentWayPoint.northSouthAxis = newNorthSouthAxis
				currentWayPoint.eastWestAxis = newEastWestAxis
			} else if instruction.value == 180 {
				newNorthSouthAxis := currentWayPoint.northSouthAxis * -1
				newEastWestAxis := currentWayPoint.eastWestAxis * -1
				currentWayPoint.northSouthAxis = newNorthSouthAxis
				currentWayPoint.eastWestAxis = newEastWestAxis
			} else if instruction.value == 270 {
				newNorthSouthAxis := currentWayPoint.eastWestAxis * -1
				newEastWestAxis := currentWayPoint.northSouthAxis
				currentWayPoint.northSouthAxis = newNorthSouthAxis
				currentWayPoint.eastWestAxis = newEastWestAxis
			}
		case right:
			if instruction.value == 90 {
				newNorthSouthAxis := currentWayPoint.eastWestAxis * -1
				newEastWestAxis := currentWayPoint.northSouthAxis
				currentWayPoint.northSouthAxis = newNorthSouthAxis
				currentWayPoint.eastWestAxis = newEastWestAxis
			} else if instruction.value == 180 {
				newNorthSouthAxis := currentWayPoint.northSouthAxis * -1
				newEastWestAxis := currentWayPoint.eastWestAxis * -1
				currentWayPoint.northSouthAxis = newNorthSouthAxis
				currentWayPoint.eastWestAxis = newEastWestAxis
			} else if instruction.value == 270 {
				newNorthSouthAxis := currentWayPoint.eastWestAxis
				newEastWestAxis := currentWayPoint.northSouthAxis * -1
				currentWayPoint.northSouthAxis = newNorthSouthAxis
				currentWayPoint.eastWestAxis = newEastWestAxis
			}
		case forward:
			northSouthAxis += instruction.value * currentWayPoint.northSouthAxis
			eastWestAxis += instruction.value * currentWayPoint.eastWestAxis
		}
	}
	fmt.Println(math.Abs(float64(northSouthAxis)) + math.Abs(float64(eastWestAxis)))
}

func main() {
	input := getInputData()
	instructions := parseData(input)
	part1(instructions)
	part2(instructions)
}

func parseData(input string) []instruction {
	instructionStrings := strings.Split(input, "\n")
	instructions := []instruction{}
	for _, val := range instructionStrings {
		var currentAction action
		switch string(val[0]) {
		case "E":
			currentAction = east
		case "W":
			currentAction = west
		case "N":
			currentAction = north
		case "S":
			currentAction = south
		case "L":
			currentAction = left
		case "R":
			currentAction = right
		case "F":
			currentAction = forward
		}
		value, _ := strconv.Atoi(val[1:])
		instructions = append(instructions, instruction{
			currentAction, value})
	}
	return instructions
}

func getInputData() string {
	data, err := ioutil.ReadFile("./input2")
	if err != nil {
		panic(err)
	}
	return string(data)
}
