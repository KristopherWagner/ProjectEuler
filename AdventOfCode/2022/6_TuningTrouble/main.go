package main

import (
	"fmt"

	"AdventOfCode/helpers"
)

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

func partOne(data string) (answer int) {
	dict := make(map[byte]int)
	dict[data[0]]++
	dict[data[1]]++
	dict[data[2]]++
	for i := 3; answer == 0 && i < len(data); i++ {
		dict[data[i]]++
		if dict[data[i-3]] == 1 && dict[data[i-2]] == 1 && dict[data[i-1]] == 1 && dict[data[i]] == 1 {
			answer = i + 1
		}
		dict[data[i-3]]--
	}
	return
}

func partTwo(data string) (answer int) {
	dict := make(map[byte]int)
	for i := 0; i < 13; i++ {
		dict[data[i]]++
	}
	for i := 13; answer == 0 && i < len(data); i++ {
		dict[data[i]]++
		valid := true
		for j := 13; valid && j >= 0; j-- {
			valid = dict[data[i-j]] == 1
		}
		if valid {
			answer = i + 1
		}
		dict[data[i-13]]--
	}
	return
}

func main() {
	/*
		fmt.Println(partOne("bvwbjplbgvbhsrlpgdmjqwftvncz") == 5)
		fmt.Println(partOne("nppdvjthqldpwncqszvftbrmjlhg") == 6)
		fmt.Println(partOne("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 10)
		fmt.Println(partOne("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 11)
	*/

	/*
		fmt.Println(partTwo("mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 19)
		fmt.Println(partTwo("bvwbjplbgvbhsrlpgdmjqwftvncz") == 23)
		fmt.Println(partTwo("nppdvjthqldpwncqszvftbrmjlhg") == 23)
		fmt.Println(partTwo("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 29)
		fmt.Println(partTwo("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 26)
	*/

	input := getInput()
	fmt.Println(partTwo(input[0]))
}
