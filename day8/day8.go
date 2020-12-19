package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := getInputData()
	instructions := parseData(input)
	visited := findVisitedInstructions(instructions)
	instructionsThatHalt := findInstructionsThatHalt(visited, instructions)
	fmt.Println(findValOfAccWhenEnteringLoop(instructionsThatHalt))

}

func findInstructionsThatHalt(visited map[int]bool, instructions []codeLine) []codeLine {
	for key, _ := range visited {
		originalInstruction := instructions[key]
		if originalInstruction.instruction == "acc" {
			continue
		}

		var modifiedOpcode string
		if originalInstruction.instruction == "nop" {
			modifiedOpcode = "jmp"
		} else {
			modifiedOpcode = "nop"
		}
		modifiedInstruction := codeLine{
			modifiedOpcode,
			originalInstruction.value}
		instructions[key] = modifiedInstruction

		if isCodeHalting(instructions) {
			return instructions
		} else {
			instructions[key] = originalInstruction
		}
	}
	return instructions
}

type codeLine struct {
	instruction string
	value       int
}

func split2(toSplit, by string) (string, string) {
	split := strings.Split(toSplit, by)
	return split[0], split[1]
}

func parseData(input string) []codeLine {
	var code []codeLine
	for _, line := range strings.Split(input, "\n") {
		instruction, valueStr := split2(line, " ")
		value, _ := strconv.Atoi(valueStr)
		code = append(code, codeLine{instruction, value})
	}
	return code
}

func isCodeHalting(instructions []codeLine) bool {
	foundLoop := false
	curLine := 0
	visited := make(map[int]bool)
	for !foundLoop && curLine < len(instructions) {
		if visited[curLine] {
			foundLoop = true
		} else {
			visited[curLine] = true
			currentCodeLine := instructions[curLine]
			switch currentCodeLine.instruction {
			case "acc":
				curLine++
			case "jmp":
				curLine += currentCodeLine.value
			case "nop":
				curLine++
			}
		}
	}
	return !foundLoop
}

func findVisitedInstructions(instructions []codeLine) map[int]bool {
	foundLoop := false
	curLine := 0
	visited := make(map[int]bool)
	for !foundLoop {
		if visited[curLine] {
			foundLoop = true
		} else {
			visited[curLine] = true
			currentCodeLine := instructions[curLine]
			switch currentCodeLine.instruction {
			case "acc":
				curLine++
				break
			case "jmp":
				curLine += currentCodeLine.value
			case "nop":
				curLine++
			}
		}
	}
	return visited
}

func findValOfAccWhenEnteringLoop(instructions []codeLine) int {
	foundLoop := false
	acc := 0
	curLine := 0
	visited := make(map[int]bool)
	for !foundLoop && curLine < len(instructions) {
		if visited[curLine] {
			foundLoop = true
		} else {
			visited[curLine] = true
			currentCodeLine := instructions[curLine]
			switch currentCodeLine.instruction {
			case "acc":
				acc += currentCodeLine.value
				curLine++
			case "jmp":
				curLine += currentCodeLine.value
			case "nop":
				curLine++
			}
		}
	}
	return acc
}

func getInputData() string {
	data, err := ioutil.ReadFile("./input2")
	if err != nil {
		panic(err)
	}
	return string(data)
}
