package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	parent *node

	left  *node
	right *node
	value int32
}

func fillNode(t *node, indexes [][]int32) {
	queue := list.New()
	current := t
	for i := 0; i < len(indexes); i++ {
		left := indexes[i][0]
		right := indexes[i][1]
		if left != -1 {
			current.left = &node{value: left}
			queue.PushBack(current.left)
		}
		if right != -1 {
			current.right = &node{value: right}
			queue.PushBack(current.right)
		}
		if queue.Len() > 0 {
			current = queue.Remove(queue.Front()).(*node)
		}
	}

}

func printTree(t *node) (indexes []int32) {
	indexes = make([]int32, 0)
	if t.left != nil {
		indexes = append(indexes, printTree(t.left)...)
	}
	indexes = append(indexes, t.value)
	if t.right != nil {
		indexes = append(indexes, printTree(t.right)...)
	}
	return
}

func performSwap(n *node) {
	if n.left == nil && n.right == nil {
		return
	}
	if n.left == nil {
		n.left = n.right
		n.right = nil
	} else if n.right == nil {
		n.right = n.left
		n.left = nil
	} else {
		temp := &node{
			value: n.left.value,
			left:  n.left.left,
			right: n.left.right,
		}
		n.left = n.right
		n.right = temp
	}
}

func swapTree(t *node, k, currentDepth int32) {
	if currentDepth%k == 0 {
		performSwap(t)
	}
	if t.left != nil {
		swapTree(t.left, k, currentDepth+1)
	}
	if t.right != nil {
		swapTree(t.right, k, currentDepth+1)
	}
}

func swapNodes(indexes [][]int32, queries []int32) (results [][]int32) {
	tree := &node{value: 1}
	fillNode(tree, indexes)
	results = make([][]int32, len(queries))
	for i, query := range queries {
		swapTree(tree, query, 1)
		results[i] = printTree(tree)
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

	var indexes [][]int32
	for i := 0; i < int(n); i++ {
		indexesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != 2 {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for i, rowItem := range result {
		for j, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if j != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

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
