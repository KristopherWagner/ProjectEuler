package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/helpers"
)

func parseOperation(line string) (first, operand, second string) {
	first = line[0:4]
	operand = string(line[5])
	second = line[7:11]
	if len(first) != 4 || len(operand) != 1 || len(second) != 4 {
		fmt.Printf("Could not parse " + line + "\n\t" + first + " " +
			operand + " " + second + "\n")
	}
	return
}

func findAnswer(values map[string]int64, operations map[string]string, name string) (answer int64) {
	answer, exists := values[name]
	if exists {
		return
	}

	operation, exists := operations[name]
	if !exists {
		fmt.Println("Could not find operation for " + name)
		return
	}
	first, operand, second := parseOperation(operation)
	firstValue := findAnswer(values, operations, first)
	secondValue := findAnswer(values, operations, second)
	switch operand {
	case "+":
		answer = firstValue + secondValue
		break
	case "-":
		answer = firstValue - secondValue
		break
	case "*":
		answer = firstValue * secondValue
		break
	case "/":
		answer = firstValue / secondValue
		break
	}

	values[name] = answer
	return
}

func partOne(data []string) (answer int64) {
	values := make(map[string]int64)
	operations := make(map[string]string)
	for i := 0; i < len(data) && len(data[i]) > 0; i++ {
		split := strings.Split(data[i], ": ")
		name := split[0]
		if number, err := strconv.ParseInt(split[1], 10, 64); err == nil {
			values[name] = number
		} else {
			operations[name] = split[1]
		}
	}

	answer = findAnswer(values, operations, "root")
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

	part1TestAnswer := int64(152)
	result := partOne(test)
	fmt.Printf("%d -> %t\n", result, result == part1TestAnswer)
	fmt.Printf("Part 1: %d\n", partOne(input))

	part2TestAnswer := int64(-1)
	result = partTwo(test)
	fmt.Printf("%d -> %t\n", result, result == part2TestAnswer)
	fmt.Printf("Part 2: %d\n", partTwo(input))
}
