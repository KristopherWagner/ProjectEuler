package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"math"
	"strconv"
)

type Node struct {
	x int
	y int

	distance int64
	value    int64
	visited  bool
}

func createGraphFromInput(input []string) (graph [][]*Node) {
	graph = make([][]*Node, len(input))
	for i := 0; i < len(input); i++ {
		line := input[i]
		graph[i] = make([]*Node, len(line))

		for j := 0; j < len(line); j++ {
			value, err := strconv.ParseInt(string(line[j]), 10, 64)
			if err == nil {
				graph[i][j] = &Node{
					x: i,
					y: j,

					distance: int64(math.MaxInt64),
					value:    value,
					visited:  false,
				}
			}
		}
	}
	return
}

func updateNeighbor(current *Node, neighbor *Node) {
	tentativeDistance := current.distance + neighbor.value
	if neighbor.distance > tentativeDistance {
		neighbor.distance = tentativeDistance
	}
	return
}

// bad implementation of Dijkstra's
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func partOne(graph [][]*Node) (result int64) {
	endI := len(graph) - 1
	endJ := len(graph[endI]) - 1

	unvisited := make([]*Node, 0)
	graph[0][0].distance = 0
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			unvisited = append(unvisited, graph[i][j])
		}
	}

	for i, j := 0, 0; len(unvisited) != 0; {
		// top
		if i != 0 && !graph[i-1][j].visited {
			updateNeighbor(graph[i][j], graph[i-1][j])
		}
		// right
		if j != len(graph[i])-1 && !graph[i][j+1].visited {
			updateNeighbor(graph[i][j], graph[i][j+1])
		}
		// bottom
		if i != len(graph)-1 && !graph[i+1][j].visited {
			updateNeighbor(graph[i][j], graph[i+1][j])
		}
		// right
		if j != 0 && !graph[i][j-1].visited {
			updateNeighbor(graph[i][j], graph[i][j-1])
		}

		graph[i][j].visited = true
		for k, found := 0, false; !found && k < len(unvisited); k++ {
			if unvisited[k] == graph[i][j] {
				unvisited = append(unvisited[:k], unvisited[k+1:]...)
				found = true
			}
		}

		for k, lowest := 0, int64(math.MaxInt64); k < len(unvisited); k++ {
			if unvisited[k].distance < lowest {
				lowest, i, j = unvisited[k].distance,
					unvisited[k].x, unvisited[k].y
			}
		}
	}

	result = graph[endI][endJ].distance
	return
}

func makeBiggerGraph(graph [][]*Node) (bigGraph [][]*Node) {
	origRows, newRows := len(graph), len(graph)*5
	origCols, newCols := len(graph[0]), len(graph[0])*5
	bigGraph = make([][]*Node, origRows*5)
	for x := 0; x < newRows; x++ {
		bigGraph[x] = make([]*Node, newCols)
		for y := 0; y < newCols; y++ {
			xi, yi := x%origRows, y%origCols
			xo, yo := x/origRows, y/origCols
			newValue := graph[xi][yi].value + int64(xo+yo)
			for newValue > 9 {
				newValue -= 9
			}
			bigGraph[x][y] = &Node{
				x: x,
				y: y,

				distance: int64(math.MaxInt64),
				value:    newValue,
				visited:  false,
			}
		}
	}
	return
}

func partTwo(graph [][]*Node) (result int64) {
	bigGraph := makeBiggerGraph(graph)
	result = partOne(bigGraph)
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	graph := createGraphFromInput(input)
	//fmt.Println(partOne(graph))
	fmt.Println(partTwo(graph))
}
