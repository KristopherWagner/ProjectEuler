package main

import "fmt"

// https://projecteuler.net/problem=14

var dictionary map[int]int

func getLength(value int) (length int) {
	length, exists := dictionary[value]
	if exists {
		return
	}

	if value == 1 {
		length = 1
		return
	}

	if value%2 == 0 {
		length = 1 + getLength(value/2)
	} else {
		length = 1 + getLength((3*value)+1)
	}

	dictionary[value] = length
	return
}

func findLongestChain(max int) (start int) {
	highest := 1
	for i := 1; i < max; i++ {
		if getLength(i) > highest {
			start = i
			highest = getLength(i)
		}
	}
	return
}

func main() {
	dictionary = make(map[int]int)
	fmt.Println(findLongestChain(1000000))
}
