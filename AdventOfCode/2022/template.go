package main

import (
	"fmt"

	"AdventOfCode/helpers"
)

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

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
		fmt.Println("Unable to ")
	}

	part1TestAnswer := 0
	fmt.Println(partOne(test) == int64(part1TestAnswer))
	fmt.Printf("Part 1: %d", partOne(input))

	part2TestAnswer := 0
	fmt.Println(partTwo(test) == int64(part2TestAnswer))
	fmt.Printf("Part 2: %d", partTwo(input))
}
