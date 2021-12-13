package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func readInput(input []string) (dots []Coordinate, maxX, maxY int, folds []string) {
	dots = make([]Coordinate, 0)
	for i := 0; i < len(input); i++ {
		line := input[i]
		split := strings.Split(line, ",")
		if len(split) == 2 {
			x, errX := strconv.ParseInt(split[0], 10, 64)
			y, errY := strconv.ParseInt(split[1], 10, 64)
			if errX == nil && errY == nil {
				dots = append(dots, Coordinate{x: int(x), y: int(y)})
				if int(x) > maxX {
					maxX = int(x)
				}
				if int(y) > maxY {
					maxY = int(y)
				}
			}
		} else if len(line) > 1 {
			folds = append(folds, line)
		}
	}
	return
}

func fillInDots(dots []Coordinate, maxX, maxY int) (paper [][]bool) {
	paper = make([][]bool, maxY)
	for i := 0; i < maxY; i++ {
		paper[i] = make([]bool, maxX)
	}

	for i := 0; i < len(dots); i++ {
		paper[dots[i].y][dots[i].x] = true
	}
	return
}

func printPaper(paper [][]bool) {
	for i := 0; i < len(paper); i++ {
		for j := 0; j < len(paper[i]); j++ {
			if paper[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
}

func parseFold(line string) (axis string, location int) {
	split := strings.Split(line, " ")
	if len(split) == 3 {
		secondSplit := strings.Split(split[2], "=")
		if len(secondSplit) == 2 {
			axis = secondSplit[0]
			location64, err := strconv.ParseInt(secondSplit[1], 10, 64)
			if err == nil {
				location = int(location64)
			}
		}
	}
	return
}

func foldPaper(paper [][]bool, axis string, location int) (newPaper [][]bool) {
	if axis == "x" {
		newPaper = make([][]bool, len(paper))
		for y := 0; y < len(paper); y++ {
			newPaper[y] = make([]bool, len(paper[y])/2)
			for x := 0; x < len(paper[y])/2; x++ {
				newPaper[y][x] = paper[y][x] || paper[y][len(paper[y])-x-1]
			}
		}
	} else if axis == "y" {
		newPaper = make([][]bool, len(paper)/2)
		for y := 0; y < len(paper)/2; y++ {
			newPaper[y] = make([]bool, len(paper[y]))
			for x := 0; x < len(paper[y]); x++ {
				newPaper[y][x] = paper[y][x] || paper[len(paper)-y-1][x]
			}
		}
	}
	return
}

func partOne(paper [][]bool, fold string) (result int) {
	axis, location := parseFold(fold)
	foldedPaper := foldPaper(paper, axis, location)
	for y := 0; y < len(foldedPaper); y++ {
		for x := 0; x < len(foldedPaper[y]); x++ {
			if foldedPaper[y][x] {
				result++
			}
		}
	}
	return
}

func partTwo(paper [][]bool, folds []string) (foldedPaper [][]bool) {
	foldedPaper = paper
	for i := 0; i < len(folds); i++ {
		axis, location := parseFold(folds[i])
		foldedPaper = foldPaper(foldedPaper, axis, location)
	}
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	dots, maxX, maxY, folds := readInput(input)
	paper := fillInDots(dots, maxX+1, maxY+1)
	//fmt.Println(partOne(paper, folds[0]))
	printPaper(partTwo(paper, folds))
}
