package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
)

func partOne() (result int64) {
	input, err := helpers.ParseInputFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dict := make(map[int]int64)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			cur, _ := strconv.ParseInt(string(input[i][j]), 10, 64)
			dict[j] += cur
		}
	}

	gammaB := ""
	epsilonB := ""
	for i := 0; i < len(input[0]); i++ {
		if dict[i] > int64(len(input)/2) {
			gammaB += "1"
			epsilonB += "0"
		} else {
			gammaB += "0"
			epsilonB += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaB, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonB, 2, 64)
	result = gamma * epsilon
	return
}

func findMostCommon(input []string, index int) (result string) {
	one := 0
	zero := 0

	for i := 0; i < len(input); i++ {
		if string(input[i][index]) == "1" {
			one++
		} else {
			zero++
		}
	}

	if one >= zero {
		result = "1"
	} else {
		result = "0"
	}
	return
}

func findOxygen(input []string, index int) (result string) {
	if len(input) == 1 {
		result = input[0]
		return
	}

	valid := make([]string, 0)

	mostCommon := findMostCommon(input, index)
	for i := 0; i < len(input); i++ {
		if string(input[i][index]) == mostCommon {
			valid = append(valid, input[i])
		}
	}

	result = findOxygen(valid, index+1)
	return
}

func findCarbon(input []string, index int) (result string) {
	if len(input) == 1 {
		result = input[0]
		return
	}

	valid := make([]string, 0)

	mostCommon := findMostCommon(input, index)
	for i := 0; i < len(input); i++ {
		if string(input[i][index]) != mostCommon {
			valid = append(valid, input[i])
		}
	}

	result = findCarbon(valid, index+1)
	return
}

func partTwo() (result int64) {
	input, err := helpers.ParseInputFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	oxygenB := findOxygen(input, 0)
	oxygen, _ := strconv.ParseInt(oxygenB, 2, 64)
	carbonB := findCarbon(input, 0)
	carbon, _ := strconv.ParseInt(carbonB, 2, 64)
	result = oxygen * carbon
	return
}

func main() {
	fmt.Println(partTwo())
}
