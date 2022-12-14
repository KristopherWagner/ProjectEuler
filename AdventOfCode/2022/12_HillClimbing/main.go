package main

import (
	"fmt"

	"AdventOfCode/helpers"
)

type Node struct {
	children []*Node
	distance int
	height   int
	visited  bool
}

var root *Node
var endNode *Node
var allNodes []*Node

func canMoveTo(start, end *Node) bool {
	dH := end.height - start.height
	return dH <= 1
}

func fillChildren(grid [][]*Node) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			start := grid[y][x]

			// left
			if x > 0 {
				end := grid[y][x-1]
				if canMoveTo(start, end) {
					start.children = append(start.children, end)
				}
			}
			// right
			if x != (len(grid[y]) - 1) {
				end := grid[y][x+1]
				if canMoveTo(start, end) {
					start.children = append(start.children, end)
				}
			}
			// up
			if y > 0 {
				end := grid[y-1][x]
				if canMoveTo(start, end) {
					start.children = append(start.children, end)
				}
			}
			// down
			if y != (len(grid) - 1) {
				end := grid[y+1][x]
				if canMoveTo(start, end) {
					start.children = append(start.children, end)
				}
			}
		}
	}
}

func fillTree(data []string) (err error) {
	grid := make([][]*Node, 0)
	for y := 0; y < len(data) && len(data[y]) > 0; y++ {
		grid = append(grid, make([]*Node, len(data[y])))
		for x := 0; x < len(data[y]); x++ {
			n := &Node{
				children: make([]*Node, 0),
				distance: 2147483647,
				height:   int(data[y][x]),
			}
			grid[y][x] = n
			if n.height == 83 {
				n.height = 97
				root = n
			} else if n.height == 69 {
				n.height = 122
				endNode = n
			}
			allNodes = append(allNodes, n)
		}
	}

	if root == nil {
		err = fmt.Errorf("did not find start")
		return
	}
	fillChildren(grid)
	return
}

func findEnd(node *Node) {
	for i := 0; i < len(node.children); i++ {
		distance := node.distance + 1
		if node.children[i].distance == -1 ||
			node.children[i].distance > distance {
			node.children[i].distance = distance
		}
	}

	node.visited = true
	var lowestNode *Node
	for i := 0; i < len(allNodes); i++ {
		if !allNodes[i].visited {
			if lowestNode == nil || allNodes[i].distance < lowestNode.distance {
				lowestNode = allNodes[i]
			}
		}
	}

	if lowestNode != nil {
		findEnd(lowestNode)
	}
}

func partOne(data []string) (answer int) {
	err := fillTree(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	root.distance = 0
	findEnd(root)
	answer = endNode.distance
	return
}

func partTwo(data []string) (answer int) {
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

	part1TestAnswer := 31
	testResult := partOne(test)
	fmt.Printf("%d -> %t\n", testResult, testResult == part1TestAnswer)
	fmt.Printf("Part 1: %d\n", partOne(input))

	part2TestAnswer := 29
	fmt.Println(partTwo(test) == part2TestAnswer)
	fmt.Printf("Part 2: %d\n", partTwo(input))
}
