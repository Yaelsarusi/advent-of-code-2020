package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := getInputData()
	numbers := parseData(input)
	sort.Ints(numbers)
	// solution := part1(numbers)
	fmt.Println(part2(numbers))
	// fmt.Println(part2(numbers, invalidNum))
}

func part2(numbers []int) int {
	pathMap := make(map[int]int)
	pathMap[0] = 1
	for _, number := range numbers {
		i := 1
		pathsToNumber := 0
		for i <= 3 {
			if val, ok := pathMap[number-i]; ok {
				pathsToNumber += val
			}
			i++
		}
		pathMap[number] = pathsToNumber
	}
	return pathMap[numbers[len(numbers)-1]]
}
func part1(numbers []int) int {
	oneJoltDiff := 0
	threeJoltDiff := 1
	previousVal := 0
	for _, val := range numbers {
		if val-previousVal == 1 {
			oneJoltDiff++
		} else if val-previousVal == 3 {
			threeJoltDiff++
		}
		previousVal = val
	}
	return oneJoltDiff * threeJoltDiff
}

func parseData(input string) []int {
	var numbers []int
	for _, str := range strings.Split(input, "\n") {
		number, _ := strconv.Atoi(str)
		numbers = append(numbers, number)
	}
	return numbers
}
func getInputData() string {
	data, err := ioutil.ReadFile("./input2")
	if err != nil {
		panic(err)
	}
	return string(data)
}
