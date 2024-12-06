package main

import (
	_ "embed"
	"strings"
)

//go:embed day4-input.txt
var input string

//go:embed day4-example.txt
var exampleInput string

func init() {
	input = strings.TrimRight(input, "\n")
	exampleInput = strings.TrimRight(exampleInput, "\n")
}

const XMAS = "XMAS"

func Part1(input string) int {
	matrix := formatInput(input)
	var result int

	var directions = map[string][2]int{
		"up":         {-1, 0},
		"down":       {1, 0},
		"left":       {0, -1},
		"right":      {0, 1},
		"up left":    {-1, -1},
		"up right":   {-1, 1},
		"down left":  {1, -1},
		"down right": {1, 1},
	}

	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] == string(XMAS[0]) {
				for d, dir := range directions {
					row, col := i+dir[0], j+dir[1]
					if dfs(matrix, row, col, 1, d, directions) {
						result += 1
					}
				}
			}
		}
	}

	return result
}

func Part2(input string) int {
	// matrix := formatInput(input)
	return 0
}

func formatInput(input string) [][]string {
	rows := strings.Split(input, "\n")
	var result [][]string
	for _, s := range rows {
		var level []string
		for _, ch := range s {
			level = append(level, string(ch))
		}
		result = append(result, level)
	}
	return result
}

func dfs(matrix [][]string, row, col, curr int, dir string, directions map[string][2]int) bool {
	if curr == len(XMAS) {
		return true
	}

	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return false
	}

	if matrix[row][col] != string(XMAS[curr]) {
		return false
	}

	r, c := row+directions[dir][0], col+directions[dir][1]

	return dfs(matrix, r, c, curr+1, dir, directions)
}
