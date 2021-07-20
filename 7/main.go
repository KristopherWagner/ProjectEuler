package main

import (
	"fmt"
	"math"
)

/*
 * https://projecteuler.net/problem=7
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
 * What is the 10 001st prime number?
 */

func isNumberPrime(value int) (isPrime bool) {
	isPrime = true
	sqrt := int(math.Sqrt(float64(value)))
	for i := 2; isPrime && i <= sqrt; i++ {
		if value%i == 0 {
			isPrime = false
		}
	}
	return
}

func findPrime(index int) (prime int) {
	numPrimes := 0
	for i := 2; numPrimes < index; i++ {
		if isNumberPrime(i) {
			prime = i
			numPrimes++
		}
	}
	return
}

func main() {
	fmt.Println(findPrime(10001))
}
