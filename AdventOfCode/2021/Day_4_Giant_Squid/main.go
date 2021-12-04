package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	numbers [5][5]int64
	marked  [5][5]bool
}

func (card *Card) generateCard(input []string) {
	for i := 0; i < 5; i++ {
		numbers := strings.Split(input[i], " ")
		for n, c := 0, 0; n < len(numbers); n++ {
			number, err := strconv.ParseInt(numbers[n], 10, 64)
			if err == nil {
				card.numbers[i][c] = number
				c++
			}
		}
	}
}

func (card *Card) handleDrawnNumber(number int64) {
	for i, found := 0, false; i < len(card.numbers) && !found; i++ {
		for j := 0; j < len(card.numbers[i]) && !found; j++ {
			if card.numbers[i][j] == number {
				card.marked[i][j], found = true, true
			}
		}
	}
}

// 0 if not a winner, a number otherwise
func (c *Card) hasWinner() (result int64) {
	isWinner := false

	// check rows
	for i := 0; i < 5 && !isWinner; i++ {
		isWinner = true
		for j := 0; j < 5 && isWinner; j++ {
			isWinner = c.marked[i][j] != false
		}
	}

	// check columns
	for i := 0; i < 5 && !isWinner; i++ {
		isWinner = true
		for j := 0; j < 5 && isWinner; j++ {
			isWinner = c.marked[j][i] != false
		}
	}

	// calculate score
	for i := 0; i < 5 && isWinner; i++ {
		for j := 0; j < 5; j++ {
			if !c.marked[i][j] {
				result += c.numbers[i][j]
			}
		}
	}

	return result
}

func cleanData(input []string) (drawn []int64, cards []Card) {
	cards = make([]Card, 0)
	drawn = make([]int64, 0)
	drawnStr := strings.Split(input[0], ",")
	for i := 0; i < len(drawnStr); i++ {
		number, err := strconv.ParseInt(drawnStr[i], 10, 64)
		if err == nil {
			drawn = append(drawn, number)
		}
	}

	for i, c := 2, 0; i < len(input); i, c = i+6, c+1 {
		card := Card{}
		card.generateCard(input[i : i+5])
		cards = append(cards, card)
	}

	return
}

func partOne(input []string) (result int64) {
	drawn, cards := cleanData(input)
	for i := 0; i < len(drawn) && result == 0; i++ {
		for c := 0; c < len(cards) && result == 0; c++ {
			cards[c].handleDrawnNumber(drawn[i])
			result = cards[c].hasWinner() * drawn[i]
		}
	}
	return
}

func partTwo(input []string) (result int64) {
	drawn, cards := cleanData(input)

	winningNumbers := make([]int64, 0)
	for i := 0; i < len(drawn); i++ {
		for c := 0; c < len(cards); c++ {
			cards[c].handleDrawnNumber(drawn[i])
			result := cards[c].hasWinner() * drawn[i]
			if result != 0 {
				winningNumbers = append(winningNumbers, result)
				cards = append(cards[:c], cards[c+1:]...)
				c--
			}
		}
	}

	result = winningNumbers[len(winningNumbers)-1]
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	//fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
