package main

import (
	"fmt"
	"math"
)

/*
 * https://projecteuler.net/problem=9
 * A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
 * a^2 + b^2 = c^2
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.
 */

// Euclids Formula
// Taken from https://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple
func generateTriple(m, n float64) (a, b, c int) {
	a = int(math.Pow(m, 2) - math.Pow(n, 2))
	b = int(2 * m * n)
	c = int(math.Pow(m, 2) + math.Pow(n, 2))
	return
}

func findProductOfSpecialTriplet() (product int) {
	for n := 1; product == 0; n++ {
		sum := 0
		for m := n + 1; sum < 1000; m++ {
			a, b, c := generateTriple(float64(m), float64(n))
			sum = a + b + c
			if sum == 1000 {
				product = a * b * c
			}
		}
	}
	return
}

func main() {
	fmt.Println(findProductOfSpecialTriplet())
}
