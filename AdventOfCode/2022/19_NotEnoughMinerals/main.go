package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"AdventOfCode/helpers"
)

type Blueprint struct {
	name          int64
	ore           int64
	clay          int64
	obsidianOre   int64
	obsidianClay  int64
	geodeOre      int64
	geodeObsidian int64
}

func findNumbersInString(line string) (numbers []int64, err error) {
	numbers = make([]int64, 0)
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	matches := re.FindAllString(line, -1)
	if len(matches) == 0 {
		return
	}

	for _, value := range matches {
		if strings.Index(value, ",") != -1 {
			value = value[:len(value)-1]
		}
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return
}

func parseBlueprint(data []string) (bp Blueprint, err error) {
	for i := 0; i < len(data) && len(data[i]) > 0; i++ {
		numbers, err := findNumbersInString(data[i])
		if err != nil {
			return bp, err
		}

		switch i {
		case 0:
			if len(numbers) != 1 {
				err = fmt.Errorf("Invalid name line: " + data[i])
				return bp, err
			}
			bp.name = numbers[0]
			break
		case 1:
			if len(numbers) != 1 {
				err = fmt.Errorf("Invalid ore line: " + data[i])
				return bp, err
			}
			bp.ore = numbers[0]
			break
		case 2:
			if len(numbers) != 1 {
				err = fmt.Errorf("Invalid clay line: " + data[i])
				return bp, err
			}
			bp.clay = numbers[0]
			break
		case 3:
			if len(numbers) != 2 {
				err = fmt.Errorf("Invalid obsidian line: " + data[i])
				return bp, err
			}
			bp.obsidianOre = numbers[0]
			bp.obsidianClay = numbers[1]
			break
		case 4:
			if len(numbers) != 2 {
				err = fmt.Errorf("Invalid geode line: " + data[i])
				return bp, err
			}
			bp.geodeOre = numbers[0]
			bp.geodeObsidian = numbers[1]
			break
		default:
			break
		}
	}
	return
}

func parseBlueprints(data []string) (bluePrints []Blueprint, err error) {
	bluePrints = make([]Blueprint, 0)
	var bp Blueprint
	for i := 0; err == nil && i < len(data); i += 6 {
		bp, err = parseBlueprint(data[i : i+5])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			bluePrints = append(bluePrints, bp)
		}
	}
	return
}

func (bp *Blueprint) simulate() (answer int64) {
	ore, oreRobots := int64(0), int64(1)
	clay, clayRobots := int64(0), int64(0)
	obsidian, obsidianRobots := int64(0), int64(0)
	geode, geodeRobots := int64(0), int64(0)
	toBuildOre, toBuildClay, toBuildObsidian, toBuildGeode := 0, 0, 0, 0

	for m := 1; m <= 24; m++ {
		fmt.Printf("\n== Minute %d ==\n", m)
		geode += geodeRobots
		obsidian += obsidianRobots
		clay += clayRobots
		ore += oreRobots

		fmt.Printf("%d ore-collecting robot collects %d ore; you now have %d ore.\n",
			oreRobots, oreRobots, ore)

		if clayRobots > 0 {
			fmt.Printf("%d clay-collecting robot collects %d clay; you now have %d clay.\n",
				clayRobots, clayRobots, clay)
		}

		if obsidianRobots > 0 {
			fmt.Printf("%d obsidian-collecting robot collects %d obsidian; you now have %d obsidian.\n",
				obsidianRobots, obsidianRobots, obsidian)
		}

		if geodeRobots > 0 {
			fmt.Printf("%d geode-collecting robot collects %d geode; you now have %d geode.\n",
				geodeRobots, geodeRobots, geode)
		}

		if toBuildGeode > 0 {
			geodeRobots++
			toBuildGeode--
		}

		if toBuildObsidian > 0 {
			obsidianRobots++
			toBuildObsidian--
		}

		if toBuildClay > 0 {
			clayRobots++
			toBuildClay--
		}

		if toBuildOre > 0 {
			oreRobots++
			toBuildOre--
		}

		if ore >= bp.geodeOre && obsidian >= bp.geodeObsidian {
			ore -= bp.geodeOre
			obsidian -= bp.geodeObsidian
			toBuildGeode++
		}

		if ore >= bp.obsidianOre && clay >= bp.obsidianClay {
			ore -= bp.obsidianOre
			clay -= bp.obsidianClay
			toBuildObsidian++
		}

		if ore >= bp.clay {
			ore -= bp.clay
			toBuildClay++
		}

		if ore >= bp.ore {
			ore -= bp.ore
			toBuildOre++
		}
	}

	fmt.Printf("Blueprint %d created %d geodes\n", bp.name, geode)
	answer = bp.name * geode
	return
}

func partOne(data []string) (answer int64) {
	bluePrints, err := parseBlueprints(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, bp := range bluePrints {
		answer += bp.simulate()
	}
	return
}

func partTwo(data []string) (answer int64) {
	return
}

func main() {
	test, err := helpers.ParseInputFile("test_data.txt")
	if err != nil {
		fmt.Println("Unable to open test file: " + err.Error())
		return
	}
	input, err := helpers.ParseInputFile("input.txt")
	if err != nil {
		fmt.Println("Unable to open input file: " + err.Error())
		return
	}

	part1TestAnswer := int64(33)
	result := partOne(test)
	fmt.Printf("%d -> %t\n", result, result == part1TestAnswer)
	//fmt.Printf("Part 1: %d\n", partOne(input))

	part2TestAnswer := int64(-1)
	result = partTwo(test)
	fmt.Printf("%d -> %t\n", result, result == part2TestAnswer)
	fmt.Printf("Part 2: %d\n", partTwo(input))
}
