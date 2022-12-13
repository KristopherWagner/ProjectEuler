package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/helpers"
)

type Point struct {
	x int64
	y int64
}

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

func getCommandValue(line string) (value int64, err error) {
	command := strings.Split(line, " ")
	if len(command) != 2 {
		err = fmt.Errorf("Invalid command" + line)
		return
	}

	value, err = strconv.ParseInt(command[1], 10, 64)
	return
}

func partOne(data []string) (answer int64) {
	x := int64(1)
	for i, cycle := 0, int64(0); i < len(data) && len(data[i]) > 0; i++ {
		if data[i][0:4] != "noop" {
			value, err := getCommandValue(data[i])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			for j := 0; j < 2; j++ {
				cycle++
				if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 ||
					cycle == 220 {
					answer += cycle * x
				}
			}
			x += value
		} else {
			cycle++
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 ||
				cycle == 220 {
				answer += cycle * x
			}
		}
	}
	return
}

func partTwo(data []string) (image [6]string) {
	spritePos := int64(1)
	cycle := 0
	for i := 0; i < len(data) && len(data[i]) > 0; i++ {
		if data[i][0:4] == "noop" {
			pixel := "."
			if int64(cycle%40) >= spritePos-1 && int64(cycle%40) <= spritePos+1 {
				pixel = "#"
			}
			image[cycle/40] += pixel
			cycle++
		} else {
			value, err := getCommandValue(data[i])
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			pixel := "."
			if int64(cycle%40) >= spritePos-1 && int64(cycle%40) <= spritePos+1 {
				pixel = "#"
			}
			image[cycle/40] += pixel
			cycle++
			pixel = "."
			if int64(cycle%40) >= spritePos-1 && int64(cycle%40) <= spritePos+1 {
				pixel = "#"
			}
			image[cycle/40] += pixel
			cycle++
			spritePos += value
		}
	}
	return
}

func main() {
	input := getInput()
	image := partTwo(input)
	for i := 0; i < 6; i++ {
		fmt.Println(image[i])
	}
}
