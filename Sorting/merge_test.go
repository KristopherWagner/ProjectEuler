package sorting

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Run("Merge sort should handle a presorted array", handlePresortedArrayMerge)
	t.Run("Merge sort should handle a reversed array", handleReversedArrayMerge)
	t.Run("Merge sort should handle a random array", handleRandomArrayMerge)
}

func handlePresortedArrayMerge(t *testing.T) {
	fmt.Printf("Testing merge sort with a presorted array")
	psA := generatePresortedArray(testArraySize)
	results, err := MergeSort(psA)
	checkTestingResults(t, results, err)
}

func handleReversedArrayMerge(t *testing.T) {
	fmt.Printf("Testing merge sort with a reversed array")
	revA := generateReversedArray(testArraySize)
	results, err := MergeSort(revA)
	checkTestingResults(t, results, err)
}

func handleRandomArrayMerge(t *testing.T) {
	var err error
	for i := 0; i < 5 && err == nil; i++ {
		fmt.Printf("Testing merge sort with a random array\t")
		randA := generateRandomArray(testArraySize)
		var results []int
		results, err = MergeSort(randA)
		checkTestingResults(t, results, err)
	}
}
