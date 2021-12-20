package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

func getBrackettedString(input string) (endIndex int, err error) {
	if string(input[0]) != "[" {
		err = fmt.Errorf("string (%s) does not start with [", input)
		return
	}
	found := false
	for i, open := 1, 1; !found && i < len(input); i++ {
		if string(input[i]) == "[" {
			open++
		} else if string(input[i]) == "]" {
			open--
			if open == 0 {
				found = true
				endIndex = i + 1
			}
		}
	}

	if !found {
		err = fmt.Errorf("could not find closing bracket in %s", input)
	}

	return
}

type Node struct {
	left  *Node // nil if a leaf
	right *Node // nil if a leaf
	value int64
}

func (n *Node) fromString(input string) (err error) {
	// [#,#]
	if len(input) == 5 {
		left, err := strconv.ParseInt(string(input[1]), 10, 64)
		if err != nil {
			return err
		}
		n.left = &Node{value: left}

		right, err := strconv.ParseInt(string(input[3]), 10, 64)
		if err != nil {
			return err
		}
		n.right = &Node{value: right}
		return err
	}

	// other cases
	if string(input[0]) != "[" {
		err = fmt.Errorf("expected [ but got " + string(input[0]))
		return
	}

	rightStr := ""
	if string(input[1]) == "[" {
		endI, err := getBrackettedString(input[1:])
		if err != nil {
			err = fmt.Errorf("could not get left string: %w", err)
			return err
		}

		leftStr := input[1 : endI+1]
		n.left = &Node{}
		n.left.fromString(leftStr)

		rightStr = input[endI+2:]
	} else {
		left, err := strconv.ParseInt(string(input[1]), 10, 64)
		if err != nil {
			err = fmt.Errorf("Could not get left from %s: %w",
				input, err)
			return err
		}
		n.left = &Node{value: left}

		rightStr = input[1:]
		rightStartIndex := strings.Index(rightStr, "[")
		if rightStartIndex == -1 {
			err = fmt.Errorf("No [ in " + rightStr)
			return err
		}
		rightStr = rightStr[rightStartIndex:]
		rightEndIndex, err := getBrackettedString(rightStr)
		if err != nil {
			err = fmt.Errorf("could not get right string: %w", err)
			return err
		}
		rightStr = rightStr[0:rightEndIndex]
	}

	if string(rightStr[0]) != "[" {
		right, err := strconv.ParseInt(string(rightStr[0]), 10, 64)
		if err != nil {
			err = fmt.Errorf("Could not get right from %s: %w",
				rightStr, err)
			return err
		}
		n.right = &Node{value: right}
	} else {
		endI, err := getBrackettedString(rightStr)
		if err != nil {
			err = fmt.Errorf("could not get right string: %w", err)
			return err
		}
		rightStr = rightStr[0:endI]
		n.right = &Node{}
		n.right.fromString(rightStr)
	}

	return
}

func (n *Node) print(prefix string) {
	if n.left != nil {
		fmt.Printf(prefix + "left-> ")
		n.left.print(prefix + "\t")
	}

	if n.right != nil {
		fmt.Printf(prefix + "right-> ")
		n.right.print(prefix + "\t")
	}

	if n.left == nil && n.right == nil {
		fmt.Printf("%d\n", n.value)
	}
}

func getNodesFromInput(input []string) (nodes []*Node) {
	nodes = make([]*Node, 0)
	for i := 0; i < len(input); i++ {
		node := &Node{}
		err := node.fromString(input[i])
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return
}

func main() {
	input, _ := helpers.ParseInputFile("test_data.txt")
	nodes := getNodesFromInput(input)
	fmt.Println(len(nodes))
}
