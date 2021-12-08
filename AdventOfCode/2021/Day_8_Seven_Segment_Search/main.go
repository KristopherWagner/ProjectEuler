package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

func parseSignalsAndOutputFromInput(input []string) (signals, output []string) {
	signals, output = make([]string, 0), make([]string, 0)
	for i := 0; i < len(input); i++ {
		split := strings.Split(input[i], " | ")
		signals = append(signals, split[0])
		output = append(output, split[1])
	}
	return
}

func partOne(output []string) (result int64) {
	for i := 0; i < len(output); i++ {
		current := strings.Split(output[i], " ")
		for j := 0; j < len(current); j++ {
			length := len(current[j])
			if length == 2 || length == 4 || length == 3 || length == 7 {
				result++
			}
		}

	}
	return
}

func findAFromOneAndSeven(one, seven string) (c string) {
	for i := 0; c == "" && i < len(seven); i++ {
		if strings.Index(one, string(seven[i])) == -1 {
			c = string(seven[i])
		}
	}
	return
}

/*
 *  aaaa
 * b    c
 * b    c
 *  dddd
 * e    f
 * e    f
 *  gggg
 */

func handleSingleDisplay(signal, output string) (result int64) {
	frequencies := make(map[string]int)
	for i := 0; i < len(signal); i++ {
		letter := string(signal[i])
		if letter != " " {
			frequencies[letter]++
		}
	}

	var one, four, seven, eight string
	numbers := strings.Split(signal, " ")
	for i := 0; i < len(numbers); i++ {
		if len(numbers[i]) == 2 {
			one = helpers.SortString(numbers[i])
		} else if len(numbers[i]) == 4 {
			four = helpers.SortString(numbers[i])
		} else if len(numbers[i]) == 3 {
			seven = helpers.SortString(numbers[i])
		} else if len(numbers[i]) == 7 {
			eight = helpers.SortString(numbers[i])
		}
	}

	var a, b, c, d, e, f, g string
	a = findAFromOneAndSeven(one, seven)

	for letter, frequency := range frequencies {
		if frequency == 6 {
			b = letter
		} else if frequency == 8 && letter != a {
			c = letter
		} else if frequency == 7 {
			if strings.Index(four, letter) != -1 {
				d = letter
			} else {
				g = letter
			}
		} else if frequency == 4 {
			e = letter
		} else if frequency == 9 {
			f = letter
		}
	}

	mapping := make(map[string]int)
	mapping[helpers.SortString(a+b+c+e+f+g)] = 0
	mapping[one] = 1
	mapping[helpers.SortString(a+c+d+e+g)] = 2
	mapping[helpers.SortString(a+c+d+f+g)] = 3
	mapping[four] = 4
	mapping[helpers.SortString(a+b+d+f+g)] = 5
	mapping[helpers.SortString(a+b+d+e+f+g)] = 6
	mapping[seven] = 7
	mapping[eight] = 8
	mapping[helpers.SortString(a+b+c+d+f+g)] = 9

	outputNumbers := strings.Split(output, " ")
	displayString := ""
	for i := 0; i < len(outputNumbers); i++ {
		current := helpers.SortString(outputNumbers[i])
		displayString = fmt.Sprintf("%s%d", displayString, mapping[current])
	}

	result, _ = strconv.ParseInt(displayString, 10, 64)
	return
}

func partTwo(signals, output []string) (result int64) {
	for i := 0; i < len(signals) && i < len(output); i++ {
		result += handleSingleDisplay(signals[i], output[i])
	}
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	signals, output := parseSignalsAndOutputFromInput(input)
	//fmt.Println(partOne(output))
	fmt.Println(partTwo(signals, output))
}
