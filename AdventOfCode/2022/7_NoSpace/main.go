package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/helpers"
)

type Node struct {
	Name string

	Children []*Node
	Parent   *Node
	Size     int64 // includes children
}

func getInput() (input []string) {
	input, _ = helpers.ParseInputFile("input.txt")
	return
}

func handleCD(cur, root *Node, directory string) (node *Node, err error) {
	if directory == "/" {
		node = cur
		return
	}

	if directory == ".." {
		if cur.Parent == nil {
			err = fmt.Errorf("parent of " + cur.Name + " is nil")
			return
		}
		node = cur.Parent
		return
	}

	if cur.Children == nil {
		err = fmt.Errorf(cur.Name + " has no children, looking for: " + directory)
		return
	}

	found := false
	for i := 0; i < len(cur.Children); i++ {
		if cur.Children[i] == nil {
			err = fmt.Errorf("Found a nil child in " + cur.Name + " looking for: " + directory)
			return
		}
		if cur.Children[i].Name == directory {
			found = true
			node = cur.Children[i]
		}
	}

	if !found {
		err = fmt.Errorf("Unable to locate " + directory + " in " + cur.Name)
	}
	return
}

func handleLS(cur *Node, output []string) (err error) {
	cur.Children = make([]*Node, len(output))
	for i := 0; i < len(output); i++ {
		split := strings.Split(output[i], " ")
		if len(split) != 2 {
			err = fmt.Errorf("Unexpected output: " + output[i])
			return
		}

		size := int64(0)
		if split[0] != "dir" {
			size, err = strconv.ParseInt(split[0], 10, 64)
			if err != nil {
				return err
			}
		}

		cur.Children[i] = &Node{
			Name:   split[1],
			Parent: cur,
			Size:   size,
		}
	}
	return
}

func generateTree(data []string) (root *Node, err error) {
	root = &Node{}
	cur := root

	for i := 0; i < len(data); i++ {
		command := data[i]

		if strings.Index(command, "$ cd") == 0 {
			split := strings.Split(command, " ")
			if len(split) != 3 {
				err = fmt.Errorf("Unexpected command: '" + command + "'\n")
				return
			}
			cur, err = handleCD(cur, root, split[2])
			if err != nil {
				return
			}
		}

		if strings.Index(command, "$ ls") == 0 {
			j := i + 1

			output := make([]string, 0)
			for found := false; !found && j < len(data) && len(data[j]) > 0; j++ {
				if string(data[j][0]) == "$" {
					found = true
				} else {
					output = append(output, data[j])
				}
			}
			err = handleLS(cur, output)
			if err != nil {
				return
			}
		}
	}

	return
}

func generateSize(node *Node) (size int64) {
	if node.Children == nil {
		size = node.Size
		return
	}

	for i := 0; i < len(node.Children); i++ {
		size += generateSize(node.Children[i])
	}
	node.Size = size
	return
}

func calculateAnswerPt1(node *Node) (answer int64) {
	if node.Children == nil {
		return
	}

	for i := 0; i < len(node.Children); i++ {
		answer += calculateAnswerPt1(node.Children[i])
	}

	if node.Size <= 100000 {
		answer += node.Size
	}
	return
}

func partOne(data []string) (answer int64, err error) {
	root, err := generateTree(data)
	if err != nil {
		return
	}
	generateSize(root)
	answer = calculateAnswerPt1(root)
	return
}

func calculateAnswerPt2(node *Node, spaceNeeded int64) (answer int64) {
	if node.Children == nil || node.Size < spaceNeeded {
		return
	}

	answer = node.Size
	for i := 0; i < len(node.Children); i++ {
		cur := calculateAnswerPt2(node.Children[i], spaceNeeded)
		if cur >= spaceNeeded && cur < answer {
			answer = cur
		}
	}
	return
}

func partTwo(data []string) (answer int64, err error) {
	root, err := generateTree(data)
	if err != nil {
		return
	}
	generateSize(root)
	answer = calculateAnswerPt2(root, 30000000-(70000000-root.Size))
	return
}

func main() {
	input := getInput()
	answer, err := partTwo(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(answer)
}
