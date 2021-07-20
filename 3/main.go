package main

import (
	"fmt"
	"math"
)

/*
 * https://projecteuler.net/problem=3
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143?
 */

/*
 * I struggled to come up with an algorithm and had to do some reading
 * I used the "Steps to find the prime factors of a number" section of the following article,
 * but did not look at their code.
 * https://www.pythonpool.com/prime-factorization-python/
 */

func findLargestPrimeFactor(value int) (largestPrimeFactor int) {
	maximum := int(math.Sqrt(float64(value)))
	for i := 2; i < maximum; i++ {
		if value%i == 0 {
			value /= i
			largestPrimeFactor = i
			i--
		}
	}
	return
}

func main() {
	fmt.Println(findLargestPrimeFactor(13195) == 29)
	fmt.Println(findLargestPrimeFactor(600851475143))
}
