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

func (p *Point) toString() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
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

func shouldMove(start, end Point) bool {
	dX := end.x - start.x
	if dX < 0 {
		dX = -dX
	}
	dY := end.y - start.y
	if dY < 0 {
		dY = -dY
	}
	return dX > 1 || dY > 1
}

func updateRope(rope []Point) []Point {
	for i := 1; i < 10 && shouldMove(rope[i], rope[i-1]); i++ {
		dx, dy := rope[i-1].x-rope[i].x, rope[i-1].y-rope[i].y
		if dx < 0 {
			rope[i].x--
		} else if dx > 0 {
			rope[i].x++
		}

		if dy < 0 {
			rope[i].y--
		} else if dy > 0 {
			rope[i].y++
		}
	}
	return rope
}

func handleMovementPt2(rope []Point, direction string, units int64, visitedLocations map[Point]bool) ([]Point, map[Point]bool) {
	for d := int64(0); d < units; d++ {
		switch direction {
		case "R":
			rope[0] = Point{x: rope[0].x + 1, y: rope[0].y}
			break
		case "U":
			rope[0] = Point{x: rope[0].x, y: rope[0].y + 1}
			break
		case "L":
			rope[0] = Point{x: rope[0].x - 1, y: rope[0].y}
			break
		case "D":
			rope[0] = Point{x: rope[0].x, y: rope[0].y - 1}
			break
		default:
			fmt.Printf("Unknown direction: " + direction)
			break
		}

		rope = updateRope(rope)
		visitedLocations[rope[9]] = true
	}
	return rope, visitedLocations
}

func partTwo(data []string) (answer int) {
	rope := make([]Point, 10)
	for i := 0; i < 10; i++ {
		rope[i] = Point{x: 0, y: 0}
	}
	visitedLocations := make(map[Point]bool)

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

	answer = len(visitedLocations)
	return
}

func main() {
	input := getInput()
	fmt.Println(partTwo(input))
}
