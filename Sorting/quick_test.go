package sorting

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("Quick sort should handle a presorted array", handlePresortedArrayQuick)
	t.Run("Quick sort should handle a reversed array", handleReversedArrayQuick)
	t.Run("Quick sort should handle a random array", handleRandomArrayQuick)
}

func handlePresortedArrayQuick(t *testing.T) {
	fmt.Printf("testing quick sort with a presorted array")
	psA := generatePresortedArray(testArraySize)
	results, err := QuickSort(psA)
	checkTestingResults(t, results, err)
}

func handleReversedArrayQuick(t *testing.T) {
	fmt.Printf("testing quick sort with a reversed array")
	revA := generateReversedArray(testArraySize)
	results, err := QuickSort(revA)
	checkTestingResults(t, results, err)
}

func handleRandomArrayQuick(t *testing.T) {
	var err error
	for i := 0; i < 5 && err == nil; i++ {
		fmt.Printf("Testing quick sort with a random array\t")
		randA := generateRandomArray(testArraySize)
		var results []int
		results, err = QuickSort(randA)
		checkTestingResults(t, results, err)
	}
}
