package main

import (
	"fmt"

	"AdventOfCode/helpers"
)

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

func isEdge(x, y int, data []string) bool {
	return y == 0 || x == 0 || y == len(data)-1 || x == len(data[y])-1
}

func isVisibleFromNorth(x, y int, data []string) (isVisible bool) {
	isVisible = true
	height := data[y][x]
	for i := y - 1; isVisible && i >= 0; i-- {
		isVisible = data[i][x] < height
	}
	return
}

func isVisibleFromSouth(x, y int, data []string) (isVisible bool) {
	isVisible = true
	height := data[y][x]
	for i := y + 1; isVisible && i < len(data); i++ {
		isVisible = data[i][x] < height
	}
	return
}

func isVisibleFromWest(x, y int, data []string) (isVisible bool) {
	isVisible = true
	height := data[y][x]
	for i := x - 1; isVisible && i >= 0; i-- {
		isVisible = data[y][i] < height
	}
	return
}

func isVisibleFromEast(x, y int, data []string) (isVisible bool) {
	isVisible = true
	height := data[y][x]
	for i := x + 1; isVisible && i < len(data[y]); i++ {
		isVisible = data[y][i] < height
	}
	return
}

func isVisible(x, y int, data []string) bool {
	if isEdge(x, y, data) {
		return true
	}

	return isVisibleFromNorth(x, y, data) ||
		isVisibleFromSouth(x, y, data) ||
		isVisibleFromWest(x, y, data) ||
		isVisibleFromEast(x, y, data)
}

func partOne(data []string) (answer int64) {
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if isVisible(x, y, data) {
				answer++
			}
		}
	}
	return
}

func calcNorth(x, y int, data []string) (score int64) {
	done := false
	for i := y - 1; i >= 0 && !done; i-- {
		score++
		done = string(data[i][x]) >= string(data[y][x])
	}
	return
}

func calcSouth(x, y int, data []string) (score int64) {
	done := false
	for i := y + 1; i < len(data) && !done; i++ {
		score++
		done = string(data[i][x]) >= string(data[y][x])
	}
	return
}

func calcEast(x, y int, data []string) (score int64) {
	done := false
	for i := x + 1; i < len(data[y]) && !done; i++ {
		score++
		done = string(data[y][i]) >= string(data[y][x])
	}
	return
}

func calcWest(x, y int, data []string) (score int64) {
	done := false
	for i := x - 1; i >= 0 && !done; i-- {
		score++
		done = string(data[y][i]) >= string(data[y][x])
	}
	return
}

func calcScore(x, y int, data []string) (score int64) {
	north := calcNorth(x, y, data)
	south := calcSouth(x, y, data)
	east := calcEast(x, y, data)
	west := calcWest(x, y, data)
	score = north * south * east * west
	return
}

func partTwo(data []string) (answer int64) {
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			score := calcScore(x, y, data)
			if score > answer {
				answer = score
			}
		}
	}
	return
}

func main() {
	input := getInput()
	fmt.Println(partTwo(input[:len(input)-1]))
}
