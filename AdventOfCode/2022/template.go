package main

import (
	"fmt"

	"AdventOfCode/helpers"
)

func partOne(data []string) (answer int64) {
	return
}

func partTwo(data []string) (answer int64) {
	return
}

func main() {
	test, err := helpers.ParseInputFile("test_data.txt")
	if err != nil {
		fmt.Println("Unable to open test file: " + err.Error())
		return
	}
	input, err := helpers.ParseInputFile("input.txt")
	if err != nil {
		fmt.Println("Unable to open input file: " + err.Error())
		return
	}

	part1TestAnswer := int64(0)
	result := partOne(test)
	fmt.Printf("%d -> %t\n", result, result == part1TestAnswer)
	fmt.Printf("Part 1: %d\n", partOne(input))

	part2TestAnswer := int64(0)
	result = partTwo(test)
	fmt.Printf("%d -> %t\n", result, result == part2TestAnswer)
	fmt.Printf("Part 2: %d\n", partTwo(input))
}
