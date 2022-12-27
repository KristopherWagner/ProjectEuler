package main

import (
	"fmt"
	"math"

	"AdventOfCode/helpers"
)

func snafuToDecimal(snafu string) (decimal int) {
	for i, pow := len(snafu)-1, 0; i >= 0; i, pow = i-1, pow+1 {
		cur := int(math.Pow(5, float64(pow)))
		switch string(snafu[i]) {
		case "=":
			decimal -= 2 * cur
			break
		case "-":
			decimal -= cur
			break
		case "0":
			break
		case "1":
			decimal += cur
			break
		case "2":
			decimal += 2 * cur
			break
		default:
			fmt.Println("Invalid character!")
			break
		}
	}
	return
}

func decimalToSnafu(decimal int) (snafu string) {
	for i := 0; decimal > 0; i++ {
		remainder := decimal % 5
		switch remainder {
		case 2:
			fallthrough
		case 1:
			fallthrough
		case 0:
			snafu = fmt.Sprintf("%d%s", remainder, snafu)
			break
		case 3:
			snafu = "=" + snafu
			decimal += 5
			break
		case 4:
			snafu = "-" + snafu
			decimal += 5
			break
		default:
			break
		}
		decimal /= 5
	}
	return
}

func partOne(data []string) (answer string) {
	sum := 0
	for i := 0; i < len(data) && len(data[i]) > 0; i++ {
		sum += snafuToDecimal(data[i])
	}
	answer = decimalToSnafu(sum)
	return
}

func runTests() bool {
	snafu := []string{
		"1",
		"2",
		"1=",
		"1-",
		"10",
		"11",
		"12",
		"2=",
		"2-",
		"20",
		"1=0",
		"1-0",
		"1=11-2",
		"1-0---0",
	}

	decimal := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		15,
		20,
		2022,
		12345,
	}

	for i := 0; i < len(snafu) && i < len(decimal); i++ {
		if snafuToDecimal(snafu[i]) != decimal[i] {
			fmt.Printf("std-%d: %d != %d\n", i, snafuToDecimal(snafu[i]), decimal[i])
			return false
		}

		if decimalToSnafu(decimal[i]) != snafu[i] {
			fmt.Printf("dts-%d: %s != %s\n", i, decimalToSnafu(decimal[i]), snafu[i])
			return false
		}
	}

	return true
}

func main() {
	if !runTests() {
		fmt.Println("Tests failed.")
		return
	}
	input, err := helpers.ParseInputFile("input.txt")
	if err != nil {
		fmt.Println("Unable to open input file: " + err.Error())
		return
	}

	fmt.Println("Part 1: " + partOne(input))
}
