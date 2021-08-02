// https://www.hackerrank.com/challenges/contacts/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// START MY CODE
type TrieNode struct {
	Children    map[rune]*TrieNode
	NumChildren int32
	IsWordEnd   bool
}

type Trie struct {
	Root *TrieNode
}

func (t *Trie) Insert(name string) {
	current := t.Root
	for _, letter := range name {
		potentialChild, exists := current.Children[letter]
		if !exists || potentialChild == nil {
			current.Children[letter] = &TrieNode{Children: make(map[rune]*TrieNode)}
		}
		current.NumChildren++
		current = current.Children[letter]
	}
	current.IsWordEnd = true
}

func (t *Trie) Find(name string) (matches int32) {
	current := t.Root
	exists := true
	for i := 0; exists && i < len(name); i++ {
		current, exists = current.Children[rune(name[i])]
	}

	if !exists {
		return
	}

	if current.IsWordEnd {
		matches = 1
	}
	matches += current.NumChildren
	return
}

func contacts(queries [][]string) (results []int32) {
	trie := &Trie{Root: &TrieNode{Children: make(map[rune]*TrieNode)}}
	// Write your code here
	for _, query := range queries {
		command := query[0]
		name := query[1]

		if command == "add" {
			trie.Insert(name)
		} else if command == "find" {
			results = append(results, trie.Find(name))
		}
	}
	return
}

// END MY CODE

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	queriesRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries [][]string
	for i := 0; i < int(queriesRows); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []string
		for _, queriesRowItem := range queriesRowTemp {
			queriesRow = append(queriesRow, queriesRowItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := contacts(queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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
