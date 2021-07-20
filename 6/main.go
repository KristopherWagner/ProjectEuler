package main

import (
	"fmt"
	"math"
)

/*
 * https://projecteuler.net/problem=6
 * There are some images with this one, so I'd just click on the link
 */

func findSumOfSquares(max int) (sum int) {
	for i := 1; i <= max; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return
}

func findSquareOfSum(max int) (square int) {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	square = int(math.Pow(float64(sum), 2))
	return
}

func main() {
	max := 100
	sumOfSquares := findSumOfSquares(max)
	squareOfSum := findSquareOfSum(max)
	fmt.Println(squareOfSum - sumOfSquares)
}
