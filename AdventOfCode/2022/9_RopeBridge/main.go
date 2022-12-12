package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"AdventOfCode/helpers"
)

type Point struct {
	x int64
	y int64
}

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("test_data.txt")
	return
}

func isAdjacent(head, tail Point) bool {
	return math.Abs(float64(head.x-tail.x)) <= 1 && math.Abs(float64(head.y-tail.y)) <= 1
}

func handleMovement(head, tail Point, direction string, units int64, visitedLocations map[Point]int64) (Point, Point, map[Point]int64) {
	visitedLocations[tail]++
	if units == 0 {
		return head, tail, visitedLocations
	}

	switch direction {
	case "R":
		head.x++
		if !isAdjacent(head, tail) {
			if head.x-tail.x > 1 {
				tail.x = head.x - 1
			}
			if math.Abs(float64(head.y-tail.y)) > 0 {
				tail.y = head.y
			}
		}
		break
	case "U":
		head.y++
		if !isAdjacent(head, tail) {
			if head.y-tail.y > 1 {
				tail.y = head.y - 1
			}
			if math.Abs(float64(head.x-tail.x)) > 0 {
				tail.x = head.x
			}
		}
		break
	case "L":
		head.x--
		if !isAdjacent(head, tail) {
			if math.Abs(float64(head.x-tail.x)) > 1 {
				tail.x = head.x + 1
			}
			if math.Abs(float64(head.y-tail.y)) > 0 {
				tail.y = head.y
			}
		}
		break
	case "D":
		head.y--
		if !isAdjacent(head, tail) {
			if math.Abs(float64(head.y-tail.y)) > 1 {
				tail.y = head.y + 1
			}
			if math.Abs(float64(head.x-tail.x)) > 0 {
				tail.x = head.x
			}
		}
		break
	default:
		fmt.Printf("Unknown direction: " + direction)
		break
	}

	return handleMovement(head, tail, direction, units-1, visitedLocations)
}

func partOne(data []string) (answer int64) {
	head := Point{x: 0, y: 0}
	tail := Point{x: 0, y: 0}
	visitedLocations := make(map[Point]int64)

	for i := 0; i < len(data) && len(data[i]) > 0; i++ {
		command := strings.Split(data[i], " ")
		if len(command) != 2 {
			fmt.Printf("Unexpected command: " + data[i] + "\n")
			return
		}

		direction := command[0]
		units, err := strconv.ParseInt(command[1], 10, 64)
		if err != nil {
			fmt.Println(err.Error() + "\n")
			return
		}
		head, tail, visitedLocations = handleMovement(head, tail, direction, units, visitedLocations)
	}

	for _, value := range visitedLocations {
		if value > 0 {
			answer++
		}
	}
	return
}

func handleMovementPt2(rope []Point, direction string, units int64, visitedLocations map[Point]int64) ([]Point, map[Point]int64) {
	// TODO
	return rope, visitedLocations
}

func partTwo(data []string) (answer int64) {
	rope := make([]Point, 10)
	for i := 0; i < 10; i++ {
		rope[i] = Point{x: 0, y: 0}
	}
	visitedLocations := make(map[Point]int64)

	for i := 0; i < len(data) && len(data[i]) > 0; i++ {
		command := strings.Split(data[i], " ")
		if len(command) != 2 {
			fmt.Printf("Unexpected command: " + data[i] + "\n")
			return
		}
		direction := command[0]
		units, err := strconv.ParseInt(command[1], 10, 64)
		if err != nil {
			fmt.Println(err.Error() + "\n")
			return
		}
		rope, visitedLocations = handleMovementPt2(rope, direction, units, visitedLocations)
	}
	for _, value := range visitedLocations {
		if value > 0 {
			answer++
		}
	}
	return
}

func main() {
	input := getInput()
	fmt.Println(partTwo(input))
}
