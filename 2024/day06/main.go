package main

import (
	_ "embed"
	"fmt"
	"strings"
	// u "github.com/BradleyRobertTaylor/advent-of-code/util"
)

//go:embed day6-input.txt
var input string

//go:embed day6-example.txt
var exampleInput string

func init() {
	input = strings.TrimRight(input, "\n")
	exampleInput = strings.TrimRight(exampleInput, "\n")
}

var directions = map[string][2]int{
	"^": {-1, 0},
	"v": {1, 0},
	">": {0, 1},
	"<": {0, -1},
}

var turns = map[string]string{
	"^": ">",
	"v": "<",
	">": "v",
	"<": "^",
}

func Part1(input string) int {
	matrix := parseInput(input)
	start := findStartPosition(matrix)
	dirKey := matrix[start[0]][start[1]]
	dir := directions[dirKey]
	visited := make(map[[2]int]bool)

	return traverse(start, matrix, dir, dirKey, [2]int{}, visited)
}

func findStartPosition(m [][]string) [2]int {
	for i, level := range m {
		for j, char := range level {
			if _, ok := directions[char]; ok {
				return [2]int{i, j}
			}
		}
	}

	return [2]int{-1, -1}
}

func traverse(start [2]int, matrix [][]string, direction [2]int, dirKey string, previousPosition [2]int, visited map[[2]int]bool) int {
	row, col := start[0], start[1]
	fmt.Println(row, col, dirKey)
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return 0
	}

	if matrix[row][col] == "#" {
		dirKey = turns[dirKey]
		direction = directions[dirKey]
		return 0 + traverse(previousPosition, matrix, direction, dirKey, [2]int{row, col}, visited)
	}

	nextPosition := [2]int{row + direction[0], col + direction[1]}
	addend := 1
	if visited[[2]int{row, col}] {
		addend = 0
	}

	visited[[2]int{row, col}] = true
	return addend + traverse(nextPosition, matrix, direction, dirKey, [2]int{row, col}, visited)
}

func Part2(input string) int {
	return 0
}

func parseInput(input string) [][]string {
	var result [][]string

	sections := strings.Split(input, "\n")
	for _, s := range sections {
		level := strings.Split(s, "")
		result = append(result, level)
	}

	return result
}
