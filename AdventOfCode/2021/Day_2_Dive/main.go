package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/helpers"
)

type Position struct {
	depth      int64
	horizontal int64
}

func parseCommand(command []string) (direction string, amt int64) {
	direction = command[0]
	amt, _ = strconv.ParseInt(command[1], 10, 64)
	return
}

func (pos *Position) handleCommand(command string) {
	direction, amt := parseCommand(strings.Split(command, " "))

	if direction == "forward" {
		pos.horizontal += amt
	} else if direction == "down" {
		pos.depth += amt
	} else if direction == "up" {
		pos.depth -= amt
	}
}

func getInput() (input []string) {
	input, err := helpers.ParseInputFile("input.txt")
	if err != nil {
		fmt.Println("Failed to get input: " + err.Error())
	}
	return
}

func partOne() (result int64) {
	input := getInput()

	pos := &Position{}
	for _, line := range input {
		pos.handleCommand(line)
	}

	return pos.horizontal * pos.depth
}

type PositionWithAim struct {
	aim int64
	Position
}

func (pos *PositionWithAim) handleCommand(command string) {
	direction, amt := parseCommand(strings.Split(command, " "))

	if direction == "forward" {
		pos.depth += pos.aim * amt
		pos.horizontal += amt
	} else if direction == "up" {
		pos.aim -= amt
	} else if direction == "down" {
		pos.aim += amt
	}
}

func partTwo() (result int64) {
	input := getInput()

	pos := &PositionWithAim{}
	for _, line := range input {
		pos.handleCommand(line)
	}

	return pos.horizontal * pos.depth
}

func main() {
	fmt.Println(partTwo())
}
