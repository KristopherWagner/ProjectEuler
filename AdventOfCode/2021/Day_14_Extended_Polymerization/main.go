package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"math"
	"strings"
)

func parseInput(input []string) (polymer string, mapping map[string]string) {
	polymer = input[0]
	mapping = make(map[string]string)
	for i := 1; i < len(input); i++ {
		split := strings.Split(input[i], " -> ")
		if len(split) == 2 {
			mapping[split[0]] = split[1]
		}
	}
	return
}

func partOne(template string, mapping map[string]string) (result int) {
	polymer := template
	for i := 0; i < 10; i++ {
		newPolymer := ""
		for j := 0; j < len(polymer)-1; j++ {
			curPair := polymer[j : j+2]
			newEle := mapping[polymer[j:j+2]]
			newPolymer += string(curPair[0]) + newEle
		}
		polymer = newPolymer + string(polymer[len(polymer)-1])
	}

	lowAmt, highAmt := len(polymer), 0
	counts := make(map[string]int)
	for i := 0; i < len(polymer); i++ {
		counts[string(polymer[i])]++
	}

	for _, v := range counts {
		if v < lowAmt {
			lowAmt = v
		}
		if v > highAmt {
			highAmt = v
		}
	}

	result = highAmt - lowAmt
	return
}

func partTwo(template string, mapping map[string]string) (result int64) {
	first := string(template[0])
	pairs := make(map[string]int64)
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for i := 0; i < 40; i++ {
		newPairs := make(map[string]int64)
		for k, v := range pairs {
			newChar := mapping[k]
			ele1, ele2 := string(k[0])+newChar, newChar+string(k[1])
			newPairs[ele1] += v
			newPairs[ele2] += v
			delete(pairs, k)
		}

		for k, v := range newPairs {
			pairs[k] = v
		}
	}

	counts := make(map[string]int64)
	for k, v := range pairs {
		counts[string(k[1])] += int64(v)
	}
	counts[first]++

	low, high := int64(math.MaxInt64), int64(0)
	for k := range counts {
		if counts[k] > high {
			high = counts[k]
		}
		if counts[k] < low {
			low = counts[k]
		}
	}

	result = high - low
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	polymer, mapping := parseInput(input)
	//fmt.Println(partOne(polymer, mapping))
	fmt.Println(partTwo(polymer, mapping))
}
