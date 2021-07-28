package sorting

import (
	"math/rand"
	"time"
)

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
