package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func getData() (data []int) {
	input, _ := helpers.ParseInputFile("input.txt")
	split := strings.Split(input[0], ",")
	data = make([]int, 0)
	for i := 0; i < len(split); i++ {
		num, err := strconv.ParseInt(split[i], 10, 64)
		if err == nil {
			data = append(data, int(num))
		}
	}
	return
}

func partOne(input []int) (result int64) {
	minFuel := int64(math.MaxInt64)
	sorted := sort.IntSlice(input)
	sort.Sort(sorted)
	for i := sorted[0]; i < sorted[len(sorted)-1]; i++ {
		curFuel := int64(0)
		for j := 0; j < len(sorted) && curFuel < minFuel; j++ {
			curFuel += int64(math.Abs(float64(sorted[j] - i)))
		}
		if curFuel < minFuel {
			minFuel = curFuel
		}
	}

	result = minFuel
	return
}

func partTwo(input []int) (result int64) {
	minFuel := int64(math.MaxInt64)
	sorted := sort.IntSlice(input)
	sort.Sort(sorted)
	for i := sorted[0]; i < sorted[len(sorted)-1]; i++ {
		curFuel := int64(0)
		for j := 0; j < len(sorted); j++ {
			distance := int64(math.Abs(float64(sorted[j] - i)))
			// I had to look up this equation...
			curFuel += (distance * (distance + 1)) / 2
		}
		if curFuel < minFuel {
			minFuel = curFuel
		}
	}
	result = minFuel
	return
}

func main() {
	input := getData()
	//fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
