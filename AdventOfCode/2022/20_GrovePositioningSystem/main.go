package main

import (
	"fmt"
	"strconv"

	"AdventOfCode/helpers"
)

type Node struct {
	value int64
	prev  *Node
	next  *Node
}

func parseData(data []string) (arr []*Node, err error) {
	for i := 0; i < len(data) && len(data[i]) > 0; i++ {
		number, err := strconv.ParseInt(data[i], 10, 64)
		if err != nil {
			return arr, err
		}
		node := &Node{value: number}
		arr = append(arr, node)
	}

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			arr[i].prev = arr[len(arr)-1]
			arr[i].next = arr[i+1]
		} else if i == len(arr)-1 {
			arr[i].prev = arr[i-1]
			arr[i].next = arr[0]
		} else {
			arr[i].prev = arr[i-1]
			arr[i].next = arr[i+1]
		}
	}
	return
}

/*
 * Before: 1st -> a -> b -> 4th
 * After: 1st -> b -> a -> 4th
 */
func swapNodes(a, b *Node) {
	first := a.prev
	last := b.next

	a.next = last
	a.prev = b

	b.next = a
	b.prev = first

	first.next = b
	last.prev = a
}

func moveNode(arr []*Node, index int) {
	current := arr[index]
	if current.value == 0 {
		return
	}

	if current.value > 0 {
		for i := 0; i < int(current.value); i++ {
			swapNodes(current, current.next)
		}
	} else if current.value < 0 {
		for i := current.value; i < 0; i++ {
			swapNodes(current.prev, current)
		}
	}
}

func findValue(start *Node, offset int) (value int64) {
	cur := start
	for i := 0; i < offset; i++ {
		cur = cur.next
	}

	value = cur.value
	return
}

func partOne(data []string) (answer int64) {
	arr, err := parseData(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < len(arr); i++ {
		moveNode(arr, i)
	}

	zeroNode := &Node{}
	found := false
	for i := 0; !found && i < len(arr); i++ {
		if arr[i].value == 0 {
			zeroNode = arr[i]
			found = true
		}
	}

	if !found {
		fmt.Println("Couldn't find zero node!")
		return
	}

	for i := 1000; i <= 3000; i += 1000 {
		answer += findValue(zeroNode, i%len(arr))
	}

	return
}

func partTwo(data []string) (answer int64) {
	return
}

func main() {
	test, err := helpers.ParseInputFile("test_data.txt")
	if err != nil {
		fmt.Println("Unable to open test file: " + err.Error())
		return
	}
	input, err := helpers.ParseInputFile("input.txt")
	if err != nil {
		fmt.Println("Unable to open input file: " + err.Error())
		return
	}

	part1TestAnswer := int64(3)
	result := partOne(test)
	fmt.Printf("%d -> %t\n", result, result == part1TestAnswer)
	fmt.Printf("Part 1: %d\n", partOne(input))

	part2TestAnswer := int64(-1)
	result = partTwo(test)
	fmt.Printf("%d -> %t\n", result, result == part2TestAnswer)
	fmt.Printf("Part 2: %d\n", partTwo(input))
}
