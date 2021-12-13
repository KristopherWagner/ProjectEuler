package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"strings"
	"unicode"
)

func addCave(caves map[string][]string, line string) (updatedCaves map[string][]string) {
	values := strings.Split(line, "-")
	first, second := values[0], values[1]

	_, exists := caves[first]
	if exists {
		caves[first] = append(caves[first], second)
	} else {
		caves[first] = []string{second}
	}

	_, exists = caves[second]
	if exists {
		caves[second] = append(caves[second], first)
	} else {
		caves[second] = []string{first}
	}

	updatedCaves = caves
	return
}

func createMap(input []string) (caves map[string][]string) {
	caves = make(map[string][]string)
	for i := 0; i < len(input); i++ {
		caves = addCave(caves, input[i])
	}
	return
}

func isUpperCase(s string) (isUpper bool) {
	isUpper = true
	for i := 0; i < len(s) && isUpper; i++ {
		r := rune(s[i])
		isUpper = unicode.IsUpper(r) && unicode.IsLetter(r)
	}
	return
}

func recurseMap(location string, caves map[string][]string, visited map[string]bool,
	pathSoFar string) (result int) {
	newVisited := make(map[string]bool)
	for k, v := range visited {
		newVisited[k] = v
	}

	if location == "end" {
		result = 1
		fmt.Println(pathSoFar)
		return
	}
	newVisited[location] = true
	connections := caves[location]
	for i := 0; i < len(connections); i++ {
		nextLocation := connections[i]
		if isUpperCase(nextLocation) || !newVisited[nextLocation] {
			result += recurseMap(nextLocation, caves, newVisited, pathSoFar+","+nextLocation)
		}
	}

	return
}

func recurseMapPt2(location string, caves map[string][]string, visited map[string]int,
	paths map[string]bool, pathSoFar string) (result map[string]bool) {
	newVisited := make(map[string]int)
	for k, v := range visited {
		newVisited[k] = v
	}

	if location == "end" {
		paths[pathSoFar] = true
		result = paths
		return
	}

	newVisited[location]--
	connections := caves[location]
	for i := 0; i < len(connections); i++ {
		nextLocation := connections[i]
		if nextLocation != "start" {
			if isUpperCase(nextLocation) || newVisited[nextLocation] > 0 {
				result = recurseMapPt2(nextLocation, caves, newVisited, paths,
					pathSoFar+","+nextLocation)
			}
		}
	}

	return
}

func partOne(caves map[string][]string) (result int) {
	visited := make(map[string]bool)
	result = recurseMap("start", caves, visited, "start")
	return
}

func partTwo(caves map[string][]string) (result int) {
	smallCaves := make([]string, 0)
	visited := make(map[string]int)
	for key := range caves {
		if !isUpperCase(key) {
			smallCaves = append(smallCaves, key)
		}
	}

	paths := make(map[string]bool)
	for i := 0; i < len(smallCaves); i++ {
		for k := range caves {
			visited[k] = 1
		}
		visited[smallCaves[i]] = 2
		paths = recurseMapPt2("start", caves, visited, paths, "start")
	}

	result = len(paths)
	return
}

func main() {
	input, _ := helpers.ParseInputFile("input.txt")
	caves := createMap(input)
	//fmt.Println(partOne(caves))
	fmt.Println(partTwo(caves))
}
