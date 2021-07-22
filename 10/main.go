package main

import (
	"fmt"
	"math"
)

/*
 * https://projecteuler.net/problem=10
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 * Find the sum of all the primes below two million.
 */

func isNumberPrime(value int) (isPrime bool) {
	isPrime = true
	sqrt := int(math.Ceil(math.Sqrt(float64(value))))
	for i := 2; isPrime && i <= sqrt; i++ {
		isPrime = value%i != 0
	}
	return
}

func findSumOfPrimes(max int) (sum int) {
	sum = 2
	for i := 3; i < max; i++ {
		if isNumberPrime(i) {
			sum += i
		}
	}
	return
}

func main() {
	fmt.Println(findSumOfPrimes(2000000))
}
