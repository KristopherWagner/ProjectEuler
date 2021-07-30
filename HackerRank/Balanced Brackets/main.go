// https://www.hackerrank.com/challenges/balanced-brackets
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
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func doCharsMatch(open, close string) (matches bool) {
	if open == "(" && close == ")" {
		matches = true
	} else if open == "{" && close == "}" {
		matches = true
	} else if open == "[" && close == "]" {
		matches = true
	}
	return
}

func isOpen(value string) bool {
	return value == "(" || value == "{" || value == "["
}

func isBalanced(s string) (balanced string) {
	balanced = "NO"
	if len(s)%2 == 0 {
		balanced = "YES"
	}

	open := ""
	for i := 0; balanced == "YES" && i < len(s); i++ {
		curLetter := string(s[i])
		if isOpen(curLetter) {
			open += curLetter
		} else {
			if len(open) > 0 {
				if doCharsMatch(string(open[len(open)-1]), curLetter) {
					open = open[0 : len(open)-1]
				} else {
					balanced = "NO"
				}
			} else {
				balanced = "NO"
			}
		}
	}

	if len(open) != 0 {
		balanced = "NO"
	}
	fmt.Println(balanced)
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Fprintf(writer, "%s\n", result)
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
