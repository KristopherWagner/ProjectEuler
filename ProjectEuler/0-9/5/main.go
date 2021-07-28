package main

import "fmt"

/*
 * https://projecteuler.net/problem=5
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any
 * remainder.
 * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */

func findSmallestNumber(maxDivisible int) (number int) {
	for i := maxDivisible; number == 0; i++ {
		valid := true
		for j := 2; j <= maxDivisible && valid; j++ {
			valid = i%j == 0
		}
		if valid {
			number = i
		}
	}

	return
}

func main() {
	fmt.Println(findSmallestNumber(20))
}
