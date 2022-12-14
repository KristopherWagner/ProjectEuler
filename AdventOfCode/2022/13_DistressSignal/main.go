package main

import (
	"fmt"

	"AdventOfCode/helpers"
)

type Packet struct {
	// TODO
	value string
}

type PacketPair struct {
	left  Packet
	right Packet
}

func (pp *PacketPair) compare() (rightOrder bool) {
	// TODO
	return
}

func partOne(data []string) (answer int64) {
	// TODO
	return
}

func partTwo(data []string) (answer int64) {
	// TODO
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

	part1TestAnswer := int64(13)
	result := partOne(test)
	fmt.Printf("%d -> %t\n", result, result == part1TestAnswer)
	fmt.Printf("Part 1: %d\n", partOne(input))

	part2TestAnswer := int64(-1)
	result = partTwo(test)
	fmt.Printf("%d -> %t\n", result, result == part2TestAnswer)
	fmt.Printf("Part 2: %d\n", partTwo(input))
}
