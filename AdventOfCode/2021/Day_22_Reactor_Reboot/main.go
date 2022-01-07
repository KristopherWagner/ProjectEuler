package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y, z int64
}

type Command struct {
	on         bool
	xmin, xmax int64
	ymin, ymax int64
	zmin, zmax int64
}

func (c *Command) isValid() (valid bool) {
	valid = true
	if c.xmin > 50 || c.xmax < -50 {
		valid = false
	}
	if c.ymin > 50 || c.ymax < -50 {
		valid = false
	}
	if c.zmin > 50 || c.zmax < -50 {
		valid = false
	}
	return
}

func (c *Command) normalize() {
	if c.xmin < -50 {
		c.xmin = -50
	}
	if c.xmax > 50 {
		c.xmax = 50
	}

	if c.ymin < -50 {
		c.ymin = -50
	}
	if c.ymax > 50 {
		c.ymax = 50
	}

	if c.zmin < -50 {
		c.zmin = -50
	}
	if c.zmax > 50 {
		c.zmax = 50
	}
}

func createReactor() (reactor map[Coordinate]bool) {
	reactor = make(map[Coordinate]bool)
	for x := int64(-50); x <= 50; x++ {
		for y := int64(-50); y <= 50; y++ {
			for z := int64(-50); z <= 50; z++ {
				reactor[Coordinate{x, y, z}] = false
			}
		}
	}
	return
}

func countLitCubes(reactor map[Coordinate]bool) (result int64) {
	for x := int64(-50); x <= 50; x++ {
		for y := int64(-50); y <= 50; y++ {
			for z := int64(-50); z <= 50; z++ {
				if reactor[Coordinate{x, y, z}] {
					result++
				}
			}
		}
	}
	return
}

func partOne(commands []Command) (result int64) {
	reactor := createReactor()
	for i := 0; i < len(commands); i++ {
		c := commands[i]
		for x := c.xmin; x <= c.xmax; x++ {
			for y := c.ymin; y <= c.ymax; y++ {
				for z := c.zmin; z <= c.zmax; z++ {
					reactor[Coordinate{x, y, z}] = c.on
				}
			}
		}
	}

	result = countLitCubes(reactor)
	return
}

func parseCommand(input string) (command Command) {
	if string(input[1]) == "n" {
		command.on = true
		input = input[3:]
	} else {
		command.on = false
		input = input[4:]
	}
	split := strings.Split(input, ",")

	xStrs := strings.Split(split[0][2:], "..")
	yStrs := strings.Split(split[1][2:], "..")
	zStrs := strings.Split(split[2][2:], "..")

	command.xmin, _ = strconv.ParseInt(xStrs[0], 10, 64)
	command.xmax, _ = strconv.ParseInt(xStrs[1], 10, 64)
	if command.xmin > command.xmax {
		command.xmin, command.xmax = command.xmax, command.xmin
	}
	command.ymin, _ = strconv.ParseInt(yStrs[0], 10, 64)
	command.ymax, _ = strconv.ParseInt(yStrs[1], 10, 64)
	if command.ymin > command.ymax {
		command.ymin, command.ymax = command.ymax, command.ymin
	}
	command.zmin, _ = strconv.ParseInt(zStrs[0], 10, 64)
	command.zmax, _ = strconv.ParseInt(zStrs[1], 10, 64)
	if command.zmin > command.zmax {
		command.zmin, command.zmax = command.zmax, command.zmin
	}
	return
}

func parseCommands(input []string) (commands []Command) {
	commands = make([]Command, 0)

	for i := 0; i < len(input); i++ {
		c := parseCommand(input[i])
		if c.isValid() {
			c.normalize()
			commands = append(commands, c)
		}
	}
	return
}

func partOneTest(fileName string, expectedValue int64) (err error) {
	input, err := helpers.ParseInputFile(fileName)
	if err != nil {
		err = fmt.Errorf("failed to open %s: %w", fileName, err)
		return
	}
	commands := parseCommands(input)
	if result := partOne(commands); result != expectedValue {
		err = fmt.Errorf("Got %d expected %d", result, expectedValue)
	}
	return
}

func main() {
	partOneTest("small_test.txt", 39)
	partOneTest("test_data.txt", 590784)

	input, _ := helpers.ParseInputFile("input.txt")
	commands := parseCommands(input)
	fmt.Println(partOne(commands))
}
