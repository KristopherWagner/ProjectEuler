package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x int64
	y int64
}

func (p *Point) fromString(str string) {
	values := strings.Split(str, ",")
	p.x, _ = strconv.ParseInt(values[0], 10, 64)
	p.y, _ = strconv.ParseInt(values[1], 10, 64)
}

func parseLine(input string) (start Point, end Point) {
	points := strings.Split(input, " -> ")
	start.fromString(points[0])
	end.fromString(points[1])
	return
}

func addLineToGrid(oldGrid [1000][1000]int64, start, end Point) (newGrid [1000][1000]int64) {
	newGrid = oldGrid
	low, high := int64(0), int64(0)

	if start.x == end.x {
		if start.y >= end.y {
			low, high = end.y, start.y
		} else {
			low, high = start.y, end.y
		}

		for i := low; i <= high; i++ {
			newGrid[start.x][i]++
		}
	}

	if start.y == end.y {
		if start.x >= end.x {
			low, high = end.x, start.x
		} else {
			low, high = start.x, end.x
		}
		for i := low; i <= high; i++ {
			newGrid[i][start.y]++
		}
	}
	return
}

func findNumOverlaps(grid [1000][1000]int64) (overlaps int64) {
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] > 1 {
				overlaps++
			}
		}
	}
	return
}

func partOne(input []string) (result int64) {
	grid := [1000][1000]int64{}

	for i := 0; i < len(input); i++ {
		line := input[i]
		if len(line) > 1 {
			start, end := parseLine(line)
			grid = addLineToGrid(grid, start, end)
		}
	}

	result = findNumOverlaps(grid)
	return
}

func getPointAlongLine(index int64, start, end Point) (p Point) {

	if start.x == end.x {
		p.x = start.x
	} else if start.x > end.x {
		p.x = start.x - index
	} else {
		p.x = start.x + index
	}

	if start.y == end.y {
		p.y = start.y
	} else if start.y > end.y {
		p.y = start.y - index
	} else {
		p.y = start.y + index
	}
	return
}

func (a *Point) equals(b Point) (areEqual bool) {
	areEqual = a.x == b.x && a.y == b.y
	return
}

func addLineToGridPartTwo(oldGrid [1000][1000]int64, start, end Point) (newGrid [1000][1000]int64) {
	newGrid = oldGrid
	for i, p := int64(1), getPointAlongLine(int64(0), start, end); !end.equals(p); i++ {
		newGrid[p.x][p.y]++
		p = getPointAlongLine(i, start, end)
	}

	newGrid[end.x][end.y]++
	return
}

func partTwo(input []string) (result int64) {
	grid := [1000][1000]int64{}
	for i := 0; i < len(input); i++ {
		line := input[i]
		if len(line) > 1 {
			start, end := parseLine(line)
			grid = addLineToGridPartTwo(grid, start, end)
		}
	}

	result = findNumOverlaps(grid)
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	//fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
