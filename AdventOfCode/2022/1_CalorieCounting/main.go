package main

import (
	"fmt"
	"strconv"

	"AdventOfCode/helpers"
)

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

func partOne(data []string) (answer int64) {
	for i, cur := 0, int64(0); i < len(data); i++ {
		if data[i] == "" {
			if cur > answer {
				answer = cur
			}
			cur = 0
		} else {
			calories, _ := strconv.ParseInt(data[i], 10, 64)
			cur += calories
		}
	}
	return
}

func partTwo(data []string) (answer int64) {
	low, mid, high := int64(0), int64(0), int64(0)
	for i, cur := 0, int64(0); i < len(data); i++ {
		if data[i] == "" {
			if cur >= high {
				low, mid, high = mid, high, cur
			} else if cur < high && cur >= mid {
				low, mid = mid, cur
			} else if cur > low && cur < mid {
				low = cur
			}
			cur = 0
		} else {
			calories, _ := strconv.ParseInt(data[i], 10, 64)
			cur += calories
		}
	}

	answer = low + mid + high
	return
}

func main() {
	input := getInput()
	fmt.Println(partTwo(input))
}
