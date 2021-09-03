// https://www.hackerrank.com/challenges/making-anagrams/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
	 * Complete the 'makeAnagram' function below.
	  *
		 * The function is expected to return an INTEGER.
		  * The function accepts following parameters:
			 *  1. STRING a
			  *  2. STRING b
*/

func makeAnagram(a string, b string) (deletions int32) {
	dictA := make(map[rune]int32)
	dictB := make(map[rune]int32)

	for _, r := range a {
		_, exists := dictA[r]
		if exists {
			dictA[r]++
		} else {
			dictA[r] = 1
		}
	}

	for _, r := range b {
		_, exists := dictB[r]
		if exists {
			dictB[r]++
		} else {
			dictB[r] = 1
		}
	}

	for i := 'a'; i <= 'z'; i++ {
		deletions += int32(math.Abs(float64(dictA[i] - dictB[i])))
	}

	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

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
