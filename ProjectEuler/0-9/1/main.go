package main

import "fmt"

/*
 * https://projecteuler.net/problem=1
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
 * The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */

func findSum(value int) (sum int) {
	for i := 0; i < value; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return
}

func main() {
	fmt.Println(findSum(10))
	fmt.Println(findSum(1000))
}
