// https://www.hackerrank.com/challenges/one-week-preparation-kit-new-year-chaos/problem
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
	 * Complete the 'minimumBribes' function below.
	  *
		 * The function accepts INTEGER_ARRAY q as parameter.
*/

func findTimesBribed(q []int32, num, pos int32) (timesBribed int32) {
	originalPosition := int32(0)
	if num-2-1 > originalPosition {
		originalPosition = num - 2 - 1
	}
	for i := pos - 1; i >= originalPosition; i-- {
		if q[i] > num {
			timesBribed++
		}
	}
	return
}

func minimumBribes(q []int32) {
	numBribes := int32(0)
	tooChaotic := false
	for i := int32(1); !tooChaotic && i <= int32(len(q)); i++ {
		if q[i-1]-i > 2 {
			tooChaotic = true
		} else {
			numBribes += findTimesBribed(q, q[i-1], i-1)
		}
	}

	if tooChaotic {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(numBribes)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
