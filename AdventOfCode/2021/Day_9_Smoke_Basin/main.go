package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
)

func isPointLow(input [][]int, i, j int) (isLowPoint bool) {
	isLowPoint = true
	point := input[i][j]

	// above
	if i > 0 && isLowPoint {
		isLowPoint = point < input[i-1][j]
	}

	// right
	if j+1 < len(input[i]) && isLowPoint {
		isLowPoint = point < input[i][j+1]
	}

	// down
	if i+1 < len(input) && isLowPoint {
		isLowPoint = point < input[i+1][j]
	}

	// left
	if j > 0 && isLowPoint {
		isLowPoint = point < input[i][j-1]
	}

	return
}

func partOne(input [][]int) (result int) {
	lowPoints := make([]int, 0)
	for i := 0; i < len(input); i++ {
		line := input[i]
		for j := 0; j < len(line); j++ {
			if isPointLow(input, i, j) {
				lowPoints = append(lowPoints, line[j])
			}
		}
	}

	for i := 0; i < len(lowPoints); i++ {
		result += lowPoints[i] + 1
	}
	return
}

func convertStringInputToIntInput(input []string) (output [][]int) {
	output = make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		line := input[i]
		output[i] = make([]int, 0)
		for j := 0; j < len(line); j++ {
			current := string(line[j])
			number, err := strconv.ParseInt(current, 10, 64)
			if err == nil {
				output[i] = append(output[i], int(number))
			}
		}
	}

	return
}

type Point struct {
	x int
	y int
}

func findLowPoints(input [][]int) (lowPoints []Point) {
	lowPoints = make([]Point, 0)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if isPointLow(input, i, j) {
				lowPoints = append(lowPoints, Point{x: i, y: j})
			}
		}
	}
	return
}

func mapBasin(input [][]int, basin map[Point]bool, lowPoint Point) {
	basin[lowPoint] = true
	i := lowPoint.x
	j := lowPoint.y
	point := input[i][j]

	// above
	if i > 0 {
		newPoint := input[i-1][j]
		if point < newPoint && newPoint != 9 {
			mapBasin(input, basin, Point{x: i - 1, y: j})
		}
	}

	// right
	if j+1 < len(input[i]) {
		newPoint := input[i][j+1]
		if point < newPoint && newPoint != 9 {
			mapBasin(input, basin, Point{x: i, y: j + 1})
		}
	}

	// down
	if i+1 < len(input) {
		newPoint := input[i+1][j]
		if point < newPoint && newPoint != 9 {
			mapBasin(input, basin, Point{x: i + 1, y: j})
		}
	}

	// left
	if j > 0 {
		newPoint := input[i][j-1]
		if point < newPoint && newPoint != 9 {
			mapBasin(input, basin, Point{x: i, y: j - 1})
		}
	}

	return
}

func partTwo(input [][]int) (result int) {
	lowPoints := findLowPoints(input)
	highest := make([]int, 3)
	for i := 0; i < len(lowPoints); i++ {
		basin := make(map[Point]bool)
		mapBasin(input, basin, lowPoints[i])

		basinSize := len(basin)
		if basinSize > highest[0] {
			highest[0], highest[1], highest[2] =
				basinSize, highest[0], highest[1]
		} else if basinSize > highest[1] {
			highest[1], highest[2] =
				basinSize, highest[1]
		} else if basinSize > highest[2] {
			highest[2] = basinSize
		}
	}

	fmt.Println(highest)
	result = highest[0] * highest[1] * highest[2]
	return
}

func main() {
	strInput, _ := helpers.ParseInputFile("input.txt")
	input := convertStringInputToIntInput(strInput)
	//fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
