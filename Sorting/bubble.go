package sorting

import "errors"

// alogirthm taken from https://en.wikipedia.org/wiki/Bubble_sort#Analysis
// best case: O(n) - when already sorted
// worse case: O(n^2) - when it's reversed
// average: O(n^2)

// BubbleSort - returns a sorted copy of values using the bubble sort algorithm
func BubbleSort(values []int) (sortedArray []int, err error) {
	sortedArray = make([]int, len(values))
	if copy(sortedArray, values) != len(values) {
		err = errors.New("failed to copy all values")
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
