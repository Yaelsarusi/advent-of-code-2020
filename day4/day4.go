package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type PassportData struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func stringToPasswordData(str string) (passportData PassportData) {
	passportData = PassportData{
		byr: "",
		iyr: "",
		eyr: "",
		hgt: "",
		hcl: "",
		ecl: "",
		pid: "",
		cid: ""}
	fields := strings.Split(strings.Replace(str, " ", "\n", -1), "\n")
	for _, field := range fields {
		fieldSplit := strings.Split(field, ":")
		switch fieldSplit[0] {
		case "byr":
			passportData.byr = fieldSplit[1]
			break
		case "iyr":
			passportData.iyr = fieldSplit[1]
			break
		case "eyr":
			passportData.eyr = fieldSplit[1]
			break
		case "hgt":
			passportData.hgt = fieldSplit[1]
			break
		case "hcl":
			passportData.hcl = fieldSplit[1]
			break
		case "ecl":
			passportData.ecl = fieldSplit[1]
			break
		case "pid":
			passportData.pid = fieldSplit[1]
			break
		case "cid":
			passportData.cid = fieldSplit[1]
			break
		}
	}
	return
}

func isStringEmpty(s string) bool {
	return len(s) == 0
}

func isByrValid(s string) bool {
	number, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return number >= 1920 && number <= 2002 && len(s) == 4
}
func isIyrValid(s string) bool {
	number, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return number >= 2010 && number <= 2020 && len(s) == 4
}
func isEyrValid(s string) bool {
	number, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return number >= 2020 && number <= 2030 && len(s) == 4
}

func isHgtValid(s string) bool {
	if len(s) > 2 {
		unit := s[len(s)-2:]
		if unit == "cm" {
			num, err := strconv.Atoi(s[:len(s)-2])
			if err != nil {
				return false
			}
			if num >= 150 && num <= 193 {
				return true
			}
		} else if unit == "in" {
			num, err := strconv.Atoi(s[:len(s)-2])
			if err != nil {
				return false
			}
			if num >= 59 && num <= 76 {
				return true
			}
		}
	}
	return false
}

func isHclValid(s string) bool {
	match, err := regexp.MatchString("^#[0-9a-f]{6}$", s)
	if err != nil {
		return false
	}
	return match

}
func isEclValid(s string) bool {
	return s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth"
}
func isPidValid(s string) bool {
	match, err := regexp.MatchString("^[0-9]{9}$", s)
	if err != nil {
		return false
	}
	return match
}

func getIsValidPassportDataPart1(data PassportData) bool {
	return !isStringEmpty(data.byr) &&
		!isStringEmpty(data.iyr) &&
		!isStringEmpty(data.eyr) &&
		!isStringEmpty(data.hgt) &&
		!isStringEmpty(data.hcl) &&
		!isStringEmpty(data.ecl) &&
		!isStringEmpty(data.pid)
}

func getIsValidPassportDataPart2(data PassportData) bool {
	return !isStringEmpty(data.byr) && isByrValid(data.byr) &&
		!isStringEmpty(data.iyr) && isIyrValid(data.iyr) &&
		!isStringEmpty(data.eyr) && isEyrValid(data.eyr) &&
		!isStringEmpty(data.hgt) && isHgtValid(data.hgt) &&
		!isStringEmpty(data.hcl) && isHclValid(data.hcl) &&
		!isStringEmpty(data.ecl) && isEclValid(data.ecl) &&
		!isStringEmpty(data.pid) && isPidValid(data.pid)
}

type getIsValidPassportData func(data PassportData) bool

func countNumberValidPassports(inputData []byte, validationFunction getIsValidPassportData) {
	validPassports := 0
	passports := strings.Split(string(inputData), "\n\n")
	for _, s := range passports {
		passportData := stringToPasswordData(s)
		if validationFunction(passportData) {
			validPassports++
		}
	}
	fmt.Println(validPassports)
}
func main() {
	dat, err := ioutil.ReadFile("./input")
	check(err)
	countNumberValidPassports(dat, getIsValidPassportDataPart1)
	countNumberValidPassports(dat, getIsValidPassportDataPart2)
}
