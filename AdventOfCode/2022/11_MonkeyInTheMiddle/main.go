package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"AdventOfCode/helpers"
)

type Monkey struct {
	number      int64
	items       []int64
	operation   string
	test        string
	ifTrue      int64
	ifFalse     int64
	inspections int64
}

func findNumberInString(line string) (number int64, err error) {
	numbers, err := findNumbersInString(line)
	if err != nil {
		return
	}

	if len(numbers) != 1 {
		err = fmt.Errorf("Found %d numbers in %s", len(numbers), line)
		return
	}
	number = numbers[0]
	return
}

func findNumbersInString(line string) (numbers []int64, err error) {
	numbers = make([]int64, 0)
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	matches := re.FindAllString(line, -1)
	if len(matches) == 0 {
		return
	}

	for _, value := range matches {
		if strings.Index(value, ",") != -1 {
			value = value[:len(value)-1]
		}
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return
}

func (m *Monkey) fillFromData(data []string) (err error) {
	if len(data) != 6 {
		err = fmt.Errorf("Length of data is %d not 6", len(data))
		return
	}
	m.number, err = findNumberInString(data[0])
	if err != nil {
		return
	}
	m.items, err = findNumbersInString(data[1])
	if err != nil {
		return
	}

	m.operation = data[2]
	m.test = data[3]

	m.ifTrue, err = findNumberInString(data[4])
	if err != nil {
		return
	}
	m.ifFalse, err = findNumberInString(data[5])
	return
}

func (m *Monkey) toString() string {
	return fmt.Sprintf("I'm Monkey %d and I have items: %+v. Inspections: %d",
		m.number, m.items, m.inspections)
}

func handleOperation(item int64, operation string) (int64, error) {
	if strings.Index(operation, "old * old") != -1 {
		return item * item, nil
	}

	if strings.Index(operation, "old + old") != -1 {
		return item + item, nil
	}

	symbol := string(operation[23])
	value, err := findNumberInString(operation)
	if err != nil {
		return item, err
	}

	switch symbol {
	case "*":
		return item * value, nil
	case "+":
		return item + value, nil
	default:
		return item, fmt.Errorf("Invalid symbol: " + symbol)
	}
}

func handleTest(item int64, test string) (bool, error) {
	if strings.Index(test, "divisible") == -1 {
		return false, fmt.Errorf("Something new! " + test)
	}
	value, err := findNumberInString(test)
	if err != nil {
		return false, err
	}

	return item%value == 0, nil
}

func simulateRound(monkeys []Monkey, divider int64) ([]Monkey, error) {
	for m := 0; m < len(monkeys); m++ {
		for len(monkeys[m].items) > 0 {
			monkeys[m].inspections++
			item := monkeys[m].items[0]
			value, err := handleOperation(item, monkeys[m].operation)
			if err != nil {
				return monkeys, err
			}

			if divider != 0 {
				value /= divider
			}

			results, err := handleTest(value, monkeys[m].test)
			if err != nil {
				return monkeys, err
			}

			var throwTo int64
			if results {
				throwTo = monkeys[m].ifTrue
			} else {
				throwTo = monkeys[m].ifFalse
			}

			monkeys[throwTo].items = append(monkeys[throwTo].items, value)
			monkeys[m].items = monkeys[m].items[1:]
		}
	}

	return monkeys, nil
}

func findAnswer(data []string, numRounds int, divider int64) (answer int64) {
	var err error
	monkeys := make([]Monkey, 0)
	for i := 0; err == nil && i < len(data)-6; i += 7 {
		m := Monkey{}
		err = m.fillFromData(data[i : i+6])
		monkeys = append(monkeys, m)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; err == nil && i < numRounds; i++ {
		monkeys, err = simulateRound(monkeys, divider)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	high := int64(0)
	higher := int64(0)
	for _, monkey := range monkeys {
		if monkey.inspections > higher {
			high, higher = higher, monkey.inspections
		} else if monkey.inspections > high {
			high = monkey.inspections
		}
	}
	answer = high * higher
	return
}

func partOne(data []string) (answer int64) {
	answer = findAnswer(data, 20, 3)
	return
}

func partTwo(data []string) (answer int64) {
	answer = findAnswer(data, 10000, 0)
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
	part1TestAnswer := 10605
	fmt.Println(partOne(test) == int64(part1TestAnswer))
	fmt.Println(partOne(input) == 182293)

	part2TestAnswer := 2713310158
	fmt.Println(partTwo(test) == int64(part2TestAnswer))
	//fmt.Printf("Part 2: %d\n", partTwo(input))
}
