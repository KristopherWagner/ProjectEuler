package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/helpers"
)

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

func getAssignments(line string) (first, second [2]int64) {
	elves := strings.Split(line, ",")
	elf := strings.Split(elves[0], "-")
	first[0], _ = strconv.ParseInt(elf[0], 10, 64)
	first[1], _ = strconv.ParseInt(elf[1], 10, 64)

	elf = strings.Split(elves[1], "-")
	second[0], _ = strconv.ParseInt(elf[0], 10, 64)
	second[1], _ = strconv.ParseInt(elf[1], 10, 64)
	return
}

func isContained(first, second [2]int64) bool {
	return second[0] >= first[0] && second[1] <= first[1]
}

func partOne(data []string) (answer int64) {
	for i := 0; i < len(data) && len(data[i]) > 6; i++ {
		first, second := getAssignments(data[i])
		if isContained(first, second) || isContained(second, first) {
			answer++
		}
	}
	return
}

func doesOverlap(first, second [2]int64) bool {
	if second[0] >= first[0] && second[0] <= first[1] {
		return true
	}
	if second[1] >= first[0] && second[1] <= first[1] {
		return true
	}
	return false
}

func partTwo(data []string) (answer int64) {
	for i := 0; i < len(data) && len(data[i]) > 6; i++ {
		first, second := getAssignments(data[i])
		if doesOverlap(first, second) || doesOverlap(second, first) {
			answer++
		}
	}
	return
}

func main() {
	input := getInput()
	fmt.Println(partTwo(input))
}
