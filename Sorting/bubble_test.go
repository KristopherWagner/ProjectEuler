package sorting

import (
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Bubble sort should handle a presorted array", handlePresortedArrayBubble)
	t.Run("Bubble sort should handle a reversed array", handleReversedArrayBubble)
	t.Run("Bubble sort should handle a random array", handleRandomArrayBubble)
}

func handlePresortedArrayBubble(t *testing.T) {
	psA := generatePresortedArray(500)
	results, err := BubbleSort(psA)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !sort.IntsAreSorted(results) {
		t.Error("Presorted array did not remain sorted")
	}
}

func handleReversedArrayBubble(t *testing.T) {
	revA := generateReversedArray(500)
	results, err := BubbleSort(revA)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !sort.IntsAreSorted(results) {
		t.Error("Failed to sort reversed array")
	}
}

func handleRandomArrayBubble(t *testing.T) {
	for i := 0; i < 5; i++ {
		randA := generateRandomArray(500)
		results, err := BubbleSort(randA)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if !sort.IntsAreSorted(results) {
			t.Error("Failed to sort random array")
		}
	}
}
