package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func groupCharectersToCountMap(groupAnswers string) (answerMap map[string]int){
	answerMap = make(map[string]int)
	for _, answer := range groupAnswers {
		answerMap[string(answer)] ++
	}
	return
}

func getNumberOfAnswersGivenByAll(answerMap map[string]int) (groupCount int){
	groupSize := answerMap["\n"] + 1
	groupCount = 0
	for _, answerCount := range answerMap {
		if answerCount == groupSize {
			groupCount ++
		}
	}
	return 
}

func getNumberOfUniqueAnswers(answerMap map[string]int) (groupCount int){
	groupCount = 0
	for answer := range answerMap {
		if answer != "\n" {
			groupCount ++
		}
	}
	return 
}

func getInputData() string {
	data, err := ioutil.ReadFile("./input")
	if (err != nil) {
		panic(err)
	}
	return string(data)
}

const groupSeperator = "\n\n"

func main() {
	groups := getInputData()
	part1Answer := 0
	part2Answer := 0
	for _, group := range strings.Split(groups, groupSeperator) {
		charToCountMap := groupCharectersToCountMap(group)
		part1Answer += getNumberOfUniqueAnswers(charToCountMap)
		part2Answer += getNumberOfAnswersGivenByAll(charToCountMap)
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}