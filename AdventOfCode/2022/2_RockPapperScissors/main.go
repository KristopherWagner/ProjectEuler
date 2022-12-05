package main

import (
	"fmt"
	"strings"

	"AdventOfCode/helpers"
)

/*
 * A, X = Rock (1)
 * B, Y = Paper (2)
 * C, Z = Scissors (3)
 */

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

func didWeDraw(mine, theirs string) (didDraw bool) {
	didDraw = (mine == "X" && theirs == "A") ||
		(mine == "Y" && theirs == "B") ||
		(mine == "Z" && theirs == "C")
	return
}

func didIWin(mine, theirs string) (didWin bool) {
	if mine == "X" {
		didWin = theirs == "C"
	} else if mine == "Y" {
		didWin = theirs == "A"
	} else if mine == "Z" {
		didWin = theirs == "B"
	}
	return
}

func calcRound(round string) (score int64) {
	choices := strings.Split(round, " ")
	mine, theirs := choices[1], choices[0]

	if mine == "X" {
		score += 1
	} else if mine == "Y" {
		score += 2
	} else if mine == "Z" {
		score += 3
	}

	if didWeDraw(mine, theirs) {
		score += 3
	} else if didIWin(mine, theirs) {
		score += 6
	}

	return
}

func partOne(data []string) (answer int64) {
	for i := 0; i < len(data) && len(data[i]) > 2; i++ {
		answer += calcRound(data[i])
	}
	return
}

/*
 * A = Rock (1)
 * B = Paper (2)
 * C = Scissors (3)
 * X = Lose (0)
 * Y = Draw (3)
 * Z = Win (6)
 */
func calcRoundPt2(round string) (score int64) {
	choices := strings.Split(round, " ")
	theirs, result := choices[0], choices[1]

	if theirs == "A" {
		if result == "X" {
			score += 0 + 3
		} else if result == "Y" {
			score += 3 + 1
		} else if result == "Z" {
			score += 6 + 2
		}
	} else if theirs == "B" {
		if result == "X" {
			score += 0 + 1
		} else if result == "Y" {
			score += 3 + 2
		} else if result == "Z" {
			score += 6 + 3
		}
	} else if theirs == "C" {
		if result == "X" {
			score += 0 + 2
		} else if result == "Y" {
			score += 3 + 3
		} else if result == "Z" {
			score += 6 + 1
		}
	}
	return
}
func partTwo(data []string) (answer int64) {
	for i := 0; i < len(data) && len(data[i]) > 2; i++ {
		answer += calcRoundPt2(data[i])
	}
	return
}

func main() {
	input := getInput()
	fmt.Println(partTwo(input))
}
