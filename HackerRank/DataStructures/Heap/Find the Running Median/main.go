package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func addToHeap(toAdd int32, heap []int32) (newHeap []int32) {
	newHeap = make([]int32, 1, 1)
	newHeap[0] = toAdd
	if len(heap) == 0 {
		return
	}
	newHeap = append(newHeap, heap...)
	for i := 0; (i+1) < len(newHeap) && newHeap[i] > newHeap[i+1]; i++ {
		newHeap[i], newHeap[i+1] = newHeap[i+1], newHeap[i]
	}
	return
}

func calculateMedian(smaller, larger []int32) (median float64) {
	if len(larger) == 0 && len(smaller) == 1 {
		median = float64(smaller[0])
		return
	}

	if (len(smaller)+len(larger))%2 == 0 {
		median = float64(smaller[len(smaller)-1]+larger[0]) / 2.0
		return
	}

	if len(smaller) > len(larger) {
		median = float64(smaller[len(smaller)-1])
		return
	}

	median = float64(larger[0])
	return
}

func rebalanceHeaps(smaller, larger []int32) (newSmaller, newLarger []int32) {
	if len(larger) > (len(smaller) + 1) {
		newSmaller = append(smaller, larger[0])
		newLarger = larger[1:]
		return
	}

	if len(smaller) > (len(larger) + 1) {
		newLarger = addToHeap(smaller[len(smaller)-1], larger)
		newSmaller = smaller[0 : len(smaller)-1]
		return
	}

	newSmaller, newLarger = smaller, larger
	return
}

func addNumberToHeaps(smaller, larger []int32, toAdd int32, lastMedian float64) (
	newSmaller, newLarger []int32) {

	if len(smaller) == 0 && len(larger) == 0 {
		newSmaller = append(smaller, toAdd)
		newLarger = larger
		return
	}

	if float64(toAdd) <= lastMedian {
		newSmaller = addToHeap(toAdd, smaller)
		newLarger = larger
	} else {
		newSmaller = smaller
		newLarger = addToHeap(toAdd, larger)
	}

	newSmaller, newLarger = rebalanceHeaps(newSmaller, newLarger)
	return
}

func runningMedian(a []int32) (medians []float64) {
	smaller := make([]int32, 0, (len(a)/2)+1)
	larger := make([]int32, 0, (len(a)/2)+1)
	medians = make([]float64, len(a), len(a))
	var lastMedian float64
	for i, number := range a {
		smaller, larger = addNumberToHeaps(smaller, larger, number, lastMedian)
		lastMedian = calculateMedian(smaller, larger)
		medians[i] = lastMedian
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	aCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var a []int32

	for i := 0; i < int(aCount); i++ {
		aItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := runningMedian(a)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%.1f", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
