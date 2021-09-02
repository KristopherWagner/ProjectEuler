// https://www.hackerrank.com/challenges/lonely-integer/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
	 * Complete the 'lonelyinteger' function below.
	  *
		 * The function is expected to return an INTEGER.
		  * The function accepts INTEGER_ARRAY a as parameter.
*/

func lonelyinteger(a []int32) (unique int32) {
	ints := make(map[int32]bool)
	for _, i := range a {
		if ints[i] {
			delete(ints, i)
		} else {
			ints[i] = true
		}
	}

	for i := range ints {
		unique = i
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := lonelyinteger(a)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}
