package sorting

// alogirthm taken from https://www.khanacademy.org/computing/computer-science/algorithms/quick-sort/a/overview-of-quicksort
// best case: O(n log n)
// worst case: O(n^2) (when the partitions are n - 1 in length)
// 	which occurs when the array is reversed in my version of the algoritm
// average: O(n log n)

func recursiveQuick(values []int) (sortedArray []int) {
	if len(values) <= 1 {
		sortedArray = values
		return
	}

	pivotIndex := len(values) - 1
	pivot := values[pivotIndex]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 0; i < pivotIndex; i++ {
		if values[i] < pivot {
			left = append(left, values[i])
		} else {
			right = append(right, values[i])
		}
	}

	left = recursiveQuick(left)
	right = recursiveQuick(right)
	sortedArray = append(left, pivot)
	sortedArray = append(sortedArray, right...)
	return
}

// QuickSort - returns a sorted copy of values using the quick sort algorithm
func QuickSort(values []int) (sortedArray []int, err error) {
	sortedArray = recursiveQuick(values)
	return
}
