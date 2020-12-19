package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func getPremable(numbers []int, start int, end int) []int {
	var premable []int
	for i := start; i < end; i++ {
		for j := i + 1; j < end; j++ {
			if numbers[i] == numbers[j] {
				continue
			}
			premable = append(premable, numbers[i]+numbers[j])
		}
	}
	return premable
}

func isNumberInPremableSums(numbers []int, number int) bool {
	for _, v := range numbers {
		if number == v {
			return true
		}
	}

	return false
}

func main() {
	input := getInputData()
	numbers := parseData(input)
	invalidNum := part1(numbers)
	fmt.Println(part2(numbers, invalidNum))
}

func part2(numbers []int, invalidNum int) int {
	startOfSet, endOfSet := 0, 0
	found := false
	currentSum := numbers[0]
	for !found {
		if currentSum == invalidNum {
			found = true
		} else if currentSum < invalidNum {
			endOfSet++
			currentSum += numbers[endOfSet]
		} else {
			currentSum -= numbers[startOfSet]
			startOfSet++
		}
	}
	sharvul := numbers[startOfSet : endOfSet+1]
	sort.Ints(sharvul)
	return sharvul[0] + sharvul[len(sharvul)-1]

}
func part1(numbers []int) int {
	for i := 25; i < len(numbers); i++ {
		sumOfPairs := getPremable(numbers, i-25, i)
		found := isNumberInPremableSums(sumOfPairs, numbers[i])
		if !found {
			return numbers[i]
		}
	}
	return -1
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
