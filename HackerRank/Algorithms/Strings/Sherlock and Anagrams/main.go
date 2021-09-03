// https://www.hackerrank.com/challenges/three-month-preparation-kit-sherlock-and-anagrams/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	 * Complete the 'sherlockAndAnagrams' function below.
	  *
		 * The function is expected to return an INTEGER.
		  * The function accepts STRING s as parameter.
*/

func sortString(unsorted string) (sorted string) {
	s := strings.Split(unsorted, "")
	sort.Strings(s)
	sorted = strings.Join(s, "")
	return
}

func hasAnagram(values []string, substring string) (anagrams int32) {
	for i := 0; i < len(values); i++ {
		if substring == values[i] {
			anagrams++
			fmt.Printf("%d %s\n", anagrams, substring)
		}
	}
	return
}

func sherlockAndAnagrams(s string) (anagrams int32) {
	dict := make(map[int][]string)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substring := sortString(s[i:j])
			value, exists := dict[len(substring)]
			if exists {
				anagrams += hasAnagram(value, substring)
			} else {
				value = make([]string, 0)
			}

			value = append(value, substring)
			dict[len(substring)] = value
		}
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
