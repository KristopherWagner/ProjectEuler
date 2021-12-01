package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() (input []string) {
	data, _ := os.ReadFile("input.txt")
	input = strings.Split(string(data), "\n")
	return
}

func cleanInput(input []string) (data []int64) {
	data = make([]int64, len(input))
	for i := 0; i < len(input); i++ {
		data[i], _ = strconv.ParseInt(input[i], 10, 64)
	}
	return
}

func partOne(data []int64) (answer int) {
	cur := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > cur {
			answer++
		}

		cur = data[i]
	}
	return
}

func calculateWindow(data []int64, startIndex int) (sum int64) {
	if startIndex+2 < len(data) {
		sum = data[startIndex] + data[startIndex+1] + data[startIndex+2]
	}
	return
}

func partTwo(data []int64) (answer int) {
	last := calculateWindow(data, 0)
	for i := 1; i < len(data)-2; i++ {
		cur := calculateWindow(data, i)
		if cur > last {
			answer++
		}
		last = cur
	}
	return
}

func main() {
	input := getInput()
	data := cleanInput(input)
	//fmt.Println(partOne(data))
	fmt.Println(partTwo(data))
}
