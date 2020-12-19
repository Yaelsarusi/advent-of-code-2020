package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func part1(minCatchTime int, buses []string) {
	minWaitTime := math.Inf(1)
	var minBus int
	for _, bus := range buses {
		if bus == "x" {
			continue
		}
		busInterval, _ := strconv.Atoi(bus)
		timeToWaitForBus := int(math.Ceil(float64(minCatchTime)/float64(busInterval))) * busInterval
		waitTime := timeToWaitForBus - minCatchTime
		if float64(waitTime) < minWaitTime {
			minWaitTime = float64(waitTime)
			minBus = busInterval
		}
	}
	fmt.Println(float64(minBus) * minWaitTime)
}

func part2(buses []string) {
	busesByMod := make(map[int]int)
	multi := 1
	for idx, bus := range buses {
		if bus == "x" {
			continue
		}
		busNum, _ := strconv.Atoi(bus)
		multi *= busNum
		busMod := (busNum - idx) % busNum
		for busMod < 0 {
			busMod += busNum
		}
		busesByMod[busNum] = busMod
	}
	toSum := 0

	for bus, mod := range busesByMod {
		sum := multi / bus
		i := 0
		for ((sum*i)%bus != mod){
			i++
		}
		toSum += sum*i
	}

	j := 0
	for (toSum - multi*j > 0 ) {
		j++
	}
	fmt.Println(toSum - multi*(j-1))
}

func main() {
	input := getInputData()
	minCatchTime, buses := parseData(input)
	part1(minCatchTime, buses)
	part2(buses)
}

func parseData(input string) (minCatchTime int, buses []string) {
	notes := strings.Split(input, "\n")
	minCatchTime, _ = strconv.Atoi(notes[0])
	buses = strings.Split(notes[1], ",")
	return
}

func getInputData() string {
	data, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	return string(data)
}
