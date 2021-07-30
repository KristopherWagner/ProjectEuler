package sorting

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Bubble sort should handle a presorted array", handlePresortedArrayBubble)
	t.Run("Bubble sort should handle a reversed array", handleReversedArrayBubble)
	t.Run("Bubble sort should handle a random array", handleRandomArrayBubble)
}

func handlePresortedArrayBubble(t *testing.T) {
	fmt.Printf("Testing bubble sort with a presorted array")
	psA := generatePresortedArray(500)
	results, err := BubbleSort(psA)
	checkTestingResults(t, results, err)
}

func handleReversedArrayBubble(t *testing.T) {
	fmt.Printf("Testing bubble sort with a reversed array")
	revA := generateReversedArray(500)
	results, err := BubbleSort(revA)
	checkTestingResults(t, results, err)
}

func handleRandomArrayBubble(t *testing.T) {
	var err error
	for i := 0; i < 5 && err == nil; i++ {
		fmt.Printf("Testing bubble sort with a random array\t")
		randA := generateRandomArray(500)
		var results []int
		results, err = BubbleSort(randA)
		checkTestingResults(t, results, err)
	}
}
