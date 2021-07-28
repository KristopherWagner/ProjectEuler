package main

import (
	"fmt"
	"math/big"
)

// https://projecteuler.net/problem=16

func calculatePowerDigit(base, power int64) (powerDigit *big.Int) {
	bigBase := big.NewInt(base)
	powerDigit = big.NewInt(base)
	for i := int64(1); i < power; i++ {
		powerDigit = powerDigit.Mul(powerDigit, bigBase)
	}
	return
}

func calculateSum(powerDigit *big.Int) (sum int64) {
	for _, char := range powerDigit.String() {
		sum += int64(char - '0')
	}
	return
}

func main() {
	powerDigit := calculatePowerDigit(2, 1000)
	fmt.Println(calculateSum(powerDigit))
}
