package main

import (
	"fmt"
	"strconv"
)

/*
 * https://projecteuler.net/problem=4
 * A palindromic number reads the same both ways.
 * The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 * Find the largest palindrome made from the product of two 3-digit numbers.
 */

func isNumberPalindrome(value int) (isPalindrome bool) {
	isPalindrome = true
	valueStr := strconv.Itoa(value)
	length := len(valueStr)

	for i := 0; i < length && isPalindrome; i++ {
		isPalindrome = valueStr[i] == valueStr[length-i-1]
	}
	return
}

func findLargestPalindrome() (largestPalindrome int) {
	largestPalindrome = 0
	for i := 999; i > 0; i-- {
		for j := 999; j > 0; j-- {
			if i*j > largestPalindrome && isNumberPalindrome(i*j) {
				largestPalindrome = i * j
			}
		}
	}
	return
}

func main() {
	fmt.Println(findLargestPalindrome())
}
