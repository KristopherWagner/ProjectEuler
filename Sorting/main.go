package sorting

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const testArraySize int = 500

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

func checkTestingResults(t *testing.T, results []int, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(results) != testArraySize || !sort.IntsAreSorted(results) {
		fmt.Printf(" - Failure\n")
		t.Error("failed to sort array")
	} else {
		fmt.Printf("\t- Success\n")
	}
}
