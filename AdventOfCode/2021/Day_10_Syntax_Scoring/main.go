package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"sort"
)

func doCharsMatch(open, close string) (matches bool) {
	if open == "(" && close == ")" {
		matches = true
	} else if open == "{" && close == "}" {
		matches = true
	} else if open == "[" && close == "]" {
		matches = true
	} else if open == "<" && close == ">" {
		matches = true
	}
	return
}

func isOpen(value string) bool {
	return value == "(" || value == "{" || value == "[" || value == "<"
}

func isBalanced(s string) (firstInvalidChar string) {
	open := ""
	for i := 0; firstInvalidChar == "" && i < len(s); i++ {
		curLetter := string(s[i])
		if isOpen(curLetter) {
			open += curLetter
		} else {
			if len(open) > 0 {
				if doCharsMatch(string(open[len(open)-1]), curLetter) {
					open = open[0 : len(open)-1]
				} else {
					firstInvalidChar = curLetter
				}
			}
		}
	}
	return
}

func calculateScore(close string) (score int64) {
	if close == ")" {
		score = 3
	} else if close == "]" {
		score = 57
	} else if close == "}" {
		score = 1197
	} else if close == ">" {
		score = 25137
	}
	return
}

func partOne(input []string) (result int64) {
	for i := 0; i < len(input); i++ {
		invalid := isBalanced(input[i])
		score := calculateScore(invalid)
		result += score
	}
	return
}

func findCompletingSymbols(s string) (result string) {
	open := ""
	for i := 0; i < len(s); i++ {
		curLetter := string(s[i])
		if isOpen(curLetter) {
			open += curLetter
		} else {
			if len(open) > 0 {
				if doCharsMatch(string(open[len(open)-1]), curLetter) {
					open = open[0 : len(open)-1]
				}
			}
		}
	}

	for i := 0; i < len(open); i++ {
		cur := string(open[i])
		if cur == "(" {
			result = ")" + result
		} else if cur == "[" {
			result = "]" + result
		} else if cur == "{" {
			result = "}" + result
		} else if cur == "<" {
			result = ">" + result
		}
	}

	return
}

func calculateScorePartTwo(closing string) (score int) {
	for i := 0; i < len(closing); i++ {
		char := string(closing[i])
		score *= 5
		if char == ")" {
			score += 1
		} else if char == "]" {
			score += 2
		} else if char == "}" {
			score += 3
		} else if char == ">" {
			score += 4
		}
	}
	return
}

func partTwo(input []string) (result int) {
	scores := make([]int, 0)
	for i := 0; i < len(input); i++ {
		invalid := isBalanced(input[i])
		if invalid == "" {
			closing := findCompletingSymbols(input[i])
			scores = append(scores, calculateScorePartTwo(closing))
		}
	}
	sorted := sort.IntSlice(scores)
	sort.Sort(sorted)
	result = sorted[len(sorted)/2]
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	//fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
