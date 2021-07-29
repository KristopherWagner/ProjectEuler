package sorting

import "fmt"

// algorithm taken from https://en.wikipedia.org/wiki/Bubble_sort#Analysis
// best case: O(n) - when already sorted
// worse case: O(n^2) - when it's reversed
// average: O(n^2)

// BubbleSort - returns a sorted copy of values using the bubble sort algorithm
func BubbleSort(values []int) (sortedArray []int, err error) {
	sortedArray, err = copyArray(values)
	if err != nil {
		err = fmt.Errorf("failed to copy array: %w", err)
		return
	}

	for hasSwapped := true; hasSwapped; {
		hasSwapped = false
		for i := 0; i < len(sortedArray)-1; i++ {
			if sortedArray[i] > sortedArray[i+1] {
				sortedArray[i], sortedArray[i+1] = sortedArray[i+1], sortedArray[i]
				hasSwapped = true
			}
		}
	}
	return
}
