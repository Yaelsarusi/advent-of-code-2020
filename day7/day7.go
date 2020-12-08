package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func checkIfBagCanContainGoldBag(bagName string, bagMap map[string]*bag) bool {
	if (bagMap[bagName].containsGoldBag) {
		return true
	}
	contains := false

	for _, containedBag := range bagMap[bagName].holdsBags {
		contains = containedBag.name == "shiny gold" || contains || checkIfBagCanContainGoldBag(containedBag.name, bagMap)
	}
	bagMap[bagName].containsGoldBag = contains

	return contains
}


func countContainedBags(bagName string, bagMap map[string]*bag) int {
	
	if (len(bagMap[bagName].holdsBags) == 0) {
		return 0
	}

	contains := 0

	for _, containedBag := range bagMap[bagName].holdsBags {
		contains += containedBag.count + containedBag.count * countContainedBags(containedBag.name, bagMap)
	}

	return contains
}

type bag struct {
	name string
	holdsBags []holdBag
	containsGoldBag bool
}

type holdBag struct {
	count int
	name string
}

func split2(toSplit, by string) (string, string){
	split := strings.Split(toSplit, by)
	return split[0], split[1]
}

func parseBags(instructions string) map[string]*bag{
	bagMap := make(map[string]*bag)

	for _, instruction := range strings.Split(instructions, "\n") {
		bagName, contains := split2(instruction, " contain")
		var containedBags []holdBag
		if contains != "" {
			containedBagsNumberAndNames := strings.Split(contains, ",")
			for _, bagNumberAndName := range containedBagsNumberAndNames {
				bagNumStr, bagName := split2(bagNumberAndName, "_")
				bagNum, _ := strconv.Atoi(bagNumStr)
				containedBags = append(containedBags, holdBag{bagNum, bagName})
			}
		}
		bagMap[bagName] = &bag{bagName, containedBags, false}
	}
	return bagMap
}

func part1(bagMap map[string]*bag) {
	canContainGoldBag := 0

	for key := range bagMap {
		if (checkIfBagCanContainGoldBag(key, bagMap)) {
			canContainGoldBag ++
		}
	}
	fmt.Println(canContainGoldBag)
}

func main() {
	bagMap := parseBags(getInputData())
	part1(bagMap)
	fmt.Println(countContainedBags("shiny gold", bagMap))
}

func getInputData() string {
	data, err := ioutil.ReadFile("./input2")
	if (err != nil) {
		panic(err)
	}
	return string(data)
}