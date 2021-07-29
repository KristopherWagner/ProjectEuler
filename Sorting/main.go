package sorting

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePresortedArray(size int) (array []int) {
	array = make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = i
	}
	return
}

func generateReversedArray(size int) (array []int) {
	array = make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = size - i
	}
	return
}

func generateRandomArray(size int) (array []int) {
	array = make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		array[i] = rand.Int()
	}
	return
}

func copyArray(toCopy []int) (theCopy []int, err error) {
	theCopy = make([]int, len(toCopy))
	numCopied := copy(theCopy, toCopy)
	if numCopied != len(toCopy) {
		err = fmt.Errorf("copied only %d of %d numbers", numCopied, len(toCopy))
	}
	return
}
