package main

import (
	"fmt"
	"math/big"
)

/*
 *  https://projecteuler.net/problem=15
 *  I used a brute force method and although I got the answer it took way too long,
 *  so I did some research.
 * 	between the following 2 articles, I found an equation that uses pascal's triangles to solve
 *  the problem.
 *  http://www.robertdickau.com/manhattan.html
 *  https://www.mathsisfun.com/pascals-triangle.html
 */

func main() {
	size := int64(20)
	var top big.Int
	top.MulRange(1, 2*size)
	var bottom big.Int
	bottom.MulRange(1, size)
	bottom.Mul(&bottom, &bottom)

	var numberOfPaths big.Int
	numberOfPaths.Div(&top, &bottom)
	fmt.Println(numberOfPaths.String())
}
