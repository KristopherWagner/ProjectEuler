package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
)

func convertStringToInt(input []string) (clean [10][10]int64) {
	for i := 0; i < len(input) && i < 10; i++ {
		for j := 0; j < len(input[i]) && j < 10; j++ {
			value, err := strconv.ParseInt(string(input[i][j]), 10, 64)
			if err == nil {
				clean[i][j] = value
			}
		}
	}

	return
}

func handleStep(input [10][10]int64) (flashes int64, newInput [10][10]int64) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			input[i][j]++
		}
	}

	for hadFlash := true; hadFlash; {
		hadFlash = false
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if input[i][j] > 9 {
					hadFlash = true
					flashes++
					input[i][j] = 0
					// top
					if i > 0 && input[i-1][j] != 0 {
						input[i-1][j]++
					}
					// top right
					if i > 0 && j < len(input)-1 && input[i-1][j+1] != 0 {
						input[i-1][j+1]++
					}
					// right
					if j < len(input)-1 && input[i][j+1] != 0 {
						input[i][j+1]++
					}
					// bottom right
					if j < len(input)-1 && i < len(input)-1 && input[i+1][j+1] != 0 {
						input[i+1][j+1]++
					}
					// bottom
					if i < len(input)-1 && input[i+1][j] != 0 {
						input[i+1][j]++
					}
					// bottom left
					if i < len(input)-1 && j > 0 && input[i+1][j-1] != 0 {
						input[i+1][j-1]++
					}
					// left
					if j > 0 && input[i][j-1] != 0 {
						input[i][j-1]++
					}
					// top left
					if i > 0 && j > 0 && input[i-1][j-1] != 0 {
						input[i-1][j-1]++
					}
				}
			}
		}
	}
	newInput = input
	return
}

func partOne(input [10][10]int64) (result int64) {
	newInput := input
	var step int64
	for i := 0; i < 100; i++ {
		step, newInput = handleStep(newInput)
		result += step
	}

	return
}

func partTwo(input [10][10]int64) (result int64) {
	newInput := input

	var step int64
	for result = 0; step != 100; step, newInput = handleStep(newInput) {
		result++
	}

	return
}

func main() {
	dirty, _ := helpers.ParseInputFile("input.txt")
	clean := convertStringToInt(dirty)
	//fmt.Println(partOne(clean))
	fmt.Println(partTwo(clean))
}
