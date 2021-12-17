package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) (xmin, xmax, ymin, ymax int64) {
	split := strings.Split(input, ": ")
	if len(split) != 2 {
		fmt.Println("Something went wrong")
		return
	}

	split = strings.Split(split[1], ", ")
	if len(split) != 2 {
		fmt.Println("Something went wrong")
		return
	}

	xstring := strings.Split(split[0][2:], "..")
	xmin, _ = strconv.ParseInt(xstring[0], 10, 64)
	xmax, _ = strconv.ParseInt(xstring[1], 10, 64)
	ystring := strings.Split(split[1][2:], "..")
	ymin, _ = strconv.ParseInt(ystring[0], 10, 64)
	ymax, _ = strconv.ParseInt(ystring[1], 10, 64)
	return
}

func testYvalue(v, ymin, ymax int64) (result int64, err error) {
	err = fmt.Errorf("did not land in zone")
	for y := int64(0); y >= ymin; v-- {
		if y > result {
			result = y
		}

		if y >= ymin && y <= ymax {
			err = nil
		}

		y += v
	}

	return
}

func partOne(ymin, ymax int64) (result int64) {
	//curValid, lastValid := false, false
	curValid := false
	//for i := int64(0); !(curValid == false && lastValid == true); i++ {
	for i := int64(0); i < 10000; i++ {
		//lastValid = curValid
		max, err := testYvalue(i, ymin, ymax)
		curValid = err == nil
		if curValid && max > result {
			result = max
		}
	}

	return
}

func testVelocity(xv, yv, xmin, xmax, ymin, ymax int64) (err error) {
	err = fmt.Errorf("did not land in zone")
	x, y := int64(0), int64(0)
	for t := int64(0); err != nil && y >= ymin && x <= xmax; t++ {
		if x >= xmin && x <= xmax && y >= ymin && y <= ymax {
			err = nil
		}

		x += xv
		y += yv
		if xv > 0 {
			xv--
		} else if xv < 0 {
			xv++
		}
		yv--
	}
	return
}

func partTwo(xmin, xmax, ymin, ymax int64) (result int64) {
	for x := int64(-1000); x < 1000; x++ {
		for y := int64(-1000); y < 1000; y++ {
			err := testVelocity(x, y, xmin, xmax, ymin, ymax)
			if err == nil {
				result++
			}
		}
	}
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	xmin, xmax, ymin, ymax := parseInput(input[0])
	//fmt.Println(partOne(ymin, ymax))
	fmt.Println(partTwo(xmin, xmax, ymin, ymax))
}
