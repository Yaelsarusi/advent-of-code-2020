package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calcRowNum(rowBinaryPartition string) int {
	min, max := 0.0, 127.0
	for _, binaryPartition := range rowBinaryPartition {
		if string(binaryPartition) == "F" {
			max = math.Floor((max + min) / 2.0)
		} else {
			min = math.Ceil((max + min) / 2.0)
		}
	}
	return int(min)
}

func calcColNum(colBinaryPartition string) int {
	min, max := 0.0, 7.0
	for _, binaryPartition := range colBinaryPartition {
		if string(binaryPartition) == "L" {
			max = math.Floor((max + min) / 2.0)
		} else {
			min = math.Ceil((max + min) / 2.0)
		}
	}
	return int(min)
}

func getSeatIDs(ticketsBinarySpacePartitioning []string) []int {
	seatIDs := make([]int, len(ticketsBinarySpacePartitioning))

	for i, ticketBinarySpacePartitioning := range ticketsBinarySpacePartitioning {
		rowBinaryPartition := ticketBinarySpacePartitioning[0:7]
		colBinaryPartition := ticketBinarySpacePartitioning[7:10]

		row := calcRowNum(rowBinaryPartition)
		col := calcColNum(colBinaryPartition)

		seatID := row*8 + col
		seatIDs[i] = seatID
	}
	return seatIDs
} 

func getNeighborsOfMissingSeatID(sortedSeatIDs []int) {
	for i := range sortedSeatIDs {
		if sortedSeatIDs[i + 1] - sortedSeatIDs[i] == 2 {
			fmt.Println(sortedSeatIDs[i], sortedSeatIDs[i + 1])
			return
		}
	}
}

func printMaxSeatID(sortedSeatIDs []int) {
	fmt.Println(sortedSeatIDs[len(sortedSeatIDs) - 1])
}
func main() {
	dat, err := ioutil.ReadFile("./input")
	check(err)
	ticketsBinarySpacePartitioning := strings.Split(string(dat), "\n")

	seatIDs := getSeatIDs(ticketsBinarySpacePartitioning)
	sort.Ints(seatIDs)
	printMaxSeatID(seatIDs)
	getNeighborsOfMissingSeatID(seatIDs)
}
