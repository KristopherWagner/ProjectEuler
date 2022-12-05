package main

import (
	"fmt"

	"AdventOfCode/helpers"
)

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

func generatePriorities() (dict map[string]int) {
	dict = make(map[string]int)
	for i := 0; i < 26; i++ {
		dict[string(rune('a'+i))] = i + 1
	}
	for i := 0; i < 26; i++ {
		dict[string(rune('A'+i))] = i + 27
	}
	return
}

func getItemInBoth(rucksack string) (item string) {
	left, right := rucksack[0:len(rucksack)/2], rucksack[len(rucksack)/2:]
	dict := make(map[string]int64)
	for i := 0; i < len(left); i++ {
		dict[string(left[i])]++
	}

	for i := 0; i < len(right) && item == ""; i++ {
		if dict[string(right[i])] > 0 {
			item = string(right[i])
		}
	}
	return
}

func partOne(data []string) (answer int64) {
	priorities := generatePriorities()
	for i := 0; i < len(data); i++ {
		item := getItemInBoth(data[i])
		answer += int64(priorities[item])
	}
	return
}

func getBadgeFromGroup(elves []string) (item string) {
	items := make(map[string]int)
	for i := 0; i < len(elves); i++ {
		current := make(map[string]int)
		for j := 0; j < len(elves[i]) && item == ""; j++ {
			letter := string(elves[i][j])

			_, exists := current[letter]
			if !exists {
				items[letter]++
				current[letter] = 1
				if items[letter] == 3 {
					item = letter
				}
			}
		}
	}
	return
}

func partTwo(data []string) (answer int64) {
	priorities := generatePriorities()
	for i := 0; i < len(data)-2; i += 3 {
		badge := getBadgeFromGroup(data[i : i+3])
		answer += int64(priorities[badge])
	}
	return
}

func main() {
	input := getInput()
	fmt.Println(partTwo(input))
}
