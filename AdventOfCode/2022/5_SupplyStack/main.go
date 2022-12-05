package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/helpers"
)

func getInstructions() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

func getInput() (input []string) {
	input = make([]string, 9)
	input[0] = "DHNQTWVB"
	input[1] = "DWB"
	input[2] = "TSQWJC"
	input[3] = "FJRNZTP"
	input[4] = "GPVJMST"
	input[5] = "BWFTN"
	input[6] = "BLDQFHVN"
	input[7] = "HPFR"
	input[8] = "ZSMBLNPH"
	return
}

func getTestInput() (input []string) {
	input = make([]string, 3)
	input[0] = "ZN"
	input[1] = "MCD"
	input[2] = "P"
	return
}

func pop(stack string) (crate, newStack string) {
	crate = string(stack[len(stack)-1])
	newStack = stack[0 : len(stack)-1]
	return
}

func handleLine(stacks []string, numToMove, from, to int64) (newStacks []string) {
	for i := 0; i < int(numToMove); i++ {
		crate, newStack := pop(stacks[from])
		newStacks = stacks
		newStacks[from] = newStack
		newStacks[to] += crate
	}
	return
}

func partOne(data []string) (answer string) {
	stacks := getInput()
	for i := 0; i < len(data) && len(data[i]) > 1; i++ {
		split := strings.Split(data[i], " ")
		numToMove, _ := strconv.ParseInt(split[1], 10, 64)
		from, _ := strconv.ParseInt(split[3], 10, 64)
		to, _ := strconv.ParseInt(split[5], 10, 64)
		stacks = handleLine(stacks, numToMove, from-1, to-1)
	}

	for i := 0; i < len(stacks); i++ {
		answer += string(stacks[i][len(stacks[i])-1])
	}
	return
}

func move(stack string, numToMove int) (crates, newStack string) {
	crates = stack[len(stack)-numToMove:]
	newStack = stack[0 : len(stack)-numToMove]
	return
}

func handleLinePt2(stacks []string, numToMove, from, to int64) (newStacks []string) {
	crates, newStack := move(stacks[from], int(numToMove))
	newStacks = stacks
	newStacks[from] = newStack
	newStacks[to] += crates
	return
}

func partTwo(data []string) (answer string) {
	stacks := getInput()
	for i := 0; i < len(data) && len(data[i]) > 1; i++ {
		split := strings.Split(data[i], " ")
		numToMove, _ := strconv.ParseInt(split[1], 10, 64)
		from, _ := strconv.ParseInt(split[3], 10, 64)
		to, _ := strconv.ParseInt(split[5], 10, 64)
		stacks = handleLinePt2(stacks, numToMove, from-1, to-1)
	}

	for i := 0; i < len(stacks); i++ {
		answer += string(stacks[i][len(stacks[i])-1])
	}
	return
}

func main() {
	input := getInstructions()
	fmt.Println(partTwo(input))
}
