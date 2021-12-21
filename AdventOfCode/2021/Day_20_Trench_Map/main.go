package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
)

func expandImage(image []string, filler string) (newImage []string) {
	newImage = make([]string, len(image)+4)

	for i := 0; i < len(image[0])+4; i++ {
		newImage[0] += filler
		newImage[1] += filler
		newImage[len(image)+2] += filler
		newImage[len(image)+3] += filler
	}

	for i := 0; i < len(image); i++ {
		newImage[i+2] = filler + filler + image[i] + filler + filler
	}

	return
}

func printImage(image []string) {
	for i := 0; i < len(image); i++ {
		fmt.Println(image[i])
	}
}

func convertPixel(image []string, algorithm, filler string, i, j int) (
	pixel string) {
	bs := ""

	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if (x >= 0 && x < len(image)) &&
				(y >= 0 && y < len(image[x])) {
				// valid
				if string(image[x][y]) == "#" {
					bs += "1"
				} else {
					bs += "0"
				}
			} else {
				// filler
				if filler == "#" {
					bs += "1"
				} else {
					bs += "0"
				}
			}
		}
	}

	bn, err := strconv.ParseInt(bs, 2, 64)
	if err != nil {
		fmt.Println(err.Error())
	}

	pixel = string(algorithm[bn])
	return
}

func applyAlogirthm(image []string, algorithm, filler string) (
	newImage []string) {
	newImage = make([]string, 0)
	for i := 0; i < len(image); i++ {
		newLine := ""
		for j := 0; j < len(image); j++ {
			newLine += convertPixel(image, algorithm, filler, i, j)
		}
		newImage = append(newImage, newLine)
	}
	return
}

func countLitPixels(image []string) (result int64) {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if string(image[i][j]) == "#" {
				result++
			}
		}
	}
	return
}

func determineFiller(lastFiller, algorithm string) (filler string) {
	if lastFiller == "." {
		filler = string(algorithm[0])
	} else {
		filler = string(algorithm[511])
	}
	return
}

func countLitPixelsAfterEnhancements(input []string,
	numEnhancements int) (result int64) {
	algorithm := input[0]
	image := input[2:]

	filler := "."

	for i := 0; i < numEnhancements; i++ {
		image = expandImage(image, filler)
		image = applyAlogirthm(image, algorithm, filler)
		filler = determineFiller(filler, algorithm)
	}
	result = countLitPixels(image)
	return
}

func main() {
	fmt.Println("Test Data")
	input1, _ := helpers.ParseInputFile("test_data.txt")
	fmt.Println(countLitPixelsAfterEnhancements(input1, 2) == 35)
	fmt.Println(countLitPixelsAfterEnhancements(input1, 50) == 3351)

	fmt.Printf("\nActual Data\n")
	input2, _ := helpers.ParseInputFile("input.txt")
	fmt.Println(countLitPixelsAfterEnhancements(input2, 2) == 5819)
	fmt.Println(countLitPixelsAfterEnhancements(input2, 50))
}
