package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	emptySeat    = "L"
	occupiedSeat = "#"
	floor        = "."
)

func getIsOccupiedSeat(rowIdx int, colIdx int, rows [][]string) bool {
	return string(rows[rowIdx][colIdx]) == occupiedSeat
}

func getSittingNeighborsPart1(rowIdx int, colIdx int, rows [][]string) int {
	neighbors := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	sittingNeighbors := 0
	for _, neighbor := range neighbors {
		if getIsOccupiedSeat(neighbor[0]+rowIdx, neighbor[1]+colIdx, rows) {
			sittingNeighbors++
		}
	}
	return sittingNeighbors
}

func getSittingNeighborsPart2(rowIdx int, colIdx int, rows [][]string) int {
	neighbors := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	sittingNeighbors := 0
	for _, neighbor := range neighbors {
		distance := 1
		currentRowIdx := neighbor[0]*distance + rowIdx
		currentColumnIndex := neighbor[1]*distance + colIdx
		for currentRowIdx > 0 && currentRowIdx < len(rows) - 1 && currentColumnIndex > 0 && currentColumnIndex < len(rows[0]) -1 && rows[currentRowIdx][currentColumnIndex] == floor {
			distance++
			currentRowIdx = neighbor[0]*distance + rowIdx
			currentColumnIndex = neighbor[1]*distance + colIdx
		}
		if getIsOccupiedSeat(currentRowIdx, currentColumnIndex, rows) {
			sittingNeighbors++
		}
	}
	return sittingNeighbors
}

func getNextRoundSeating(currentSeat string, sittingNeighbors int) string {
	if currentSeat == emptySeat && sittingNeighbors == 0 {
		return occupiedSeat
	}
	if currentSeat == occupiedSeat && sittingNeighbors >= 5 {
		return emptySeat
	}
	if currentSeat == "" || currentSeat == floor {
		return floor
	}
	return currentSeat

}

func getSeatsAndHasChanged(lastSeats [][]string) ([][]string, bool) {
	hasChanged := false
	input := getInputData()
	seats := parseData(input)

	for rowIdx := 1; rowIdx < len(lastSeats)-1; rowIdx++ {
		for colIdx := 1; colIdx < len(lastSeats[0])-1; colIdx++ {
			sittingNeighbors := getSittingNeighborsPart2(rowIdx, colIdx, lastSeats)

			nextRoundSeat := getNextRoundSeating(string(lastSeats[rowIdx][colIdx]), sittingNeighbors)

			if nextRoundSeat != string(lastSeats[rowIdx][colIdx]) {
				hasChanged = true
			}
			seats[rowIdx][colIdx] = nextRoundSeat
		}
	}
	return seats, hasChanged
}

func main() {
	input := getInputData()
	lastSeats := parseData(input)
	hasChanged := true

	for hasChanged {
		lastSeats, hasChanged = getSeatsAndHasChanged(lastSeats)
	}
	sitting := 0

	for rowIdx := 1; rowIdx < len(lastSeats)-1; rowIdx++ {
		for colIdx := 1; colIdx < len(lastSeats[0])-1; colIdx++ {
			if getIsOccupiedSeat(rowIdx, colIdx, lastSeats) {
				sitting++
			}
		}
	}
	fmt.Println(sitting)
}

func parseData(input string) [][]string {
	rows := strings.Split(input, "\n")
	seats := make([][]string, 97)
	for i := range seats {
		seats[i] = make([]string, 101)
	}

	for idx, row := range rows {
		seats[idx] = strings.Split(string(row), "")
	}
	return seats
}

func getInputData() string {
	data, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	return string(data)
}
