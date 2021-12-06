package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Lanternfish struct {
	timer int64
}

func (lf *Lanternfish) init() {
	lf.timer = 8
}

func (lf *Lanternfish) tick() (didReproduce bool) {
	if lf.timer == 0 {
		didReproduce = true
		lf.timer = 6
	} else {
		lf.timer--
	}
	return
}

func calculateFish(input string, days int) (result int) {
	initialValues := strings.Split(input, ",")
	fish := make([]Lanternfish, 0)

	for i := 0; i < len(initialValues); i++ {
		value, err := strconv.ParseInt(initialValues[i], 10, 64)
		if err == nil {
			lf := Lanternfish{timer: value}
			fish = append(fish, lf)
		}
	}

	for i := 0; i < days; i++ {
		newFish := make([]Lanternfish, 0)
		for j := 0; j < len(fish); j++ {
			if fish[j].tick() {
				baby := Lanternfish{}
				baby.init()
				newFish = append(newFish, baby)
			}
		}
		fish = append(fish, newFish...)
	}

	result = len(fish)
	return
}

func partOne(input string) (result int) {
	result = calculateFish(input, 80)
	return
}

func partTwo(input string, days int) (result int64) {
	initialFish := strings.Split(input, ",")
	fish := make([]int64, 9) // index = days until reproduction

	for i := 0; i < len(initialFish); i++ {
		value, err := strconv.ParseInt(initialFish[i], 10, 64)
		if err == nil {
			fish[value]++
		}
	}

	for i := 0; i < days; i++ {
		numBabies := fish[0]
		for j := 0; j < 8; j++ {
			fish[j] = fish[j+1] // decrement each timer
		}
		fish[6] += numBabies // fish tht reproduced
		fish[8] = numBabies  // the actual babies
	}

	for i := 0; i < 9; i++ {
		result += fish[i]
	}
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	//fmt.Println(partOne(input[0]))
	fmt.Println(partTwo(input[0], 256))
}
