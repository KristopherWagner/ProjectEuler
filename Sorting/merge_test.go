package sorting

import (
	"fmt"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Run("Merge sort should handle a presorted array", handlePresortedArrayMerge)
	t.Run("Merge sort should handle a reversed array", handleReversedArrayMerge)
	t.Run("Merge sort should handle a random array", handleRandomArrayMerge)
}

func handlePresortedArrayMerge(t *testing.T) {
	fmt.Printf("Testing merge sort with a presorted array")
	psA := generatePresortedArray(500)
	results, err := MergeSort(psA)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !sort.IntsAreSorted(results) {
		t.Error("Presorted array did not remain sorted")
	} else {
		fmt.Printf(" - Success\n")
	}
}

func handleReversedArrayMerge(t *testing.T) {
	fmt.Printf("Testing merge sort with a reversed array")
	revA := generateReversedArray(500)
	results, err := MergeSort(revA)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !sort.IntsAreSorted(results) {
		t.Error("Failed to sort reversed array")
	} else {
		fmt.Printf(" - Success\n")
	}
}

func handleRandomArrayMerge(t *testing.T) {
	fmt.Printf("Testing merge sort with multiple random arrays")
	var err error
	for i := 0; i < 5 && err == nil; i++ {
		randA := generateRandomArray(500)
		var results []int
		results, err = MergeSort(randA)
		if err != nil {
			fmt.Printf(" - Failure\n")
			t.Fatalf(err.Error())
		}
		if !sort.IntsAreSorted(results) {
			fmt.Printf(" - Failure\n")
			err = fmt.Errorf("failed to sort random array")
			t.Error(err.Error())
		}
	}
	if err == nil {
		fmt.Printf(" - Success\n")
	}
}
