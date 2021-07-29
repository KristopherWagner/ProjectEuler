package sorting

import (
	"fmt"
)

// algorithm https://en.wikipedia.org/wiki/Merge_sort#Algorithm
// best case: O(n log n) (but half as much time as worst case)
// worst case: O(n log n)
// average: O(n log n)
// One downside is that the array is not sorted in place and uses a bunch of memory

func merge(left, right []int) (merged []int) {
	mergedSize := len(left) + len(right)
	merged = make([]int, mergedSize)

	for m, l, r := 0, 0, 0; m < mergedSize; m++ {
		if l == len(left) {
			merged[m] = right[r]
			r++
		} else if r == len(right) {
			merged[m] = left[l]
			l++
		} else {
			if left[l] < right[r] {
				merged[m] = left[l]
				l++
			} else {
				merged[m] = right[r]
				r++
			}
		}
	}
	return
}

func recursiveMerge(values []int) (sortedValues []int, err error) {
	if len(values) <= 1 {
		sortedValues = values
		return
	}

	middle := len(values) / 2
	left, err := recursiveMerge(values[0:middle])
	if err != nil {
		err = fmt.Errorf("failed to sort left array: %w", err)
		return
	}
	right, err := recursiveMerge(values[middle:])
	if err != nil {
		err = fmt.Errorf("failed to sort left array: %w", err)
		return
	}
	sortedValues = merge(left, right)
	return
}

// MergeSort - returns a sorted copy of values using the merge sort algorithm
func MergeSort(values []int) (sortedArray []int, err error) {
	sortedArray, err = recursiveMerge(values)
	if err != nil {
		err = fmt.Errorf("failed to sort array using merge sort: %w", err)
	}
	return
}
