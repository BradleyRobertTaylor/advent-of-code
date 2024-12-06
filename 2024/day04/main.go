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
const MAS = "MAS"

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
						result++
					}
				}
			}
		}
	}

	return result
}

func Part2(input string) int {
	matrix := formatInput(input)
	var result int

	for i := 0; i < len(matrix)-2; i++ {
		for j := 0; j < len(matrix[0])-2; j++ {
			if isCrossMAS(matrix, i, j) {
				result++
			}
		}
	}

	return result
}

func isCrossMAS(matrix [][]string, row int, col int) bool {
	corners := map[string][2]int{
		"leftUp":    {row, col},
		"rightUp":   {row, col + 2},
		"leftDown":  {row + 2, col},
		"rightDown": {row + 2, col + 2},
	}

	var count int
	for k, v := range corners {
		isMAS := true
		switch k {
		case "leftUp":
			for i := range len(MAS) {
				r, c := v[0]+i, v[1]+i
				if matrix[r][c] != string(MAS[i]) {
					isMAS = false
				}
			}
		case "rightUp":
			for i := range len(MAS) {
				r, c := v[0]+i, v[1]-i
				if matrix[r][c] != string(MAS[i]) {
					isMAS = false
				}
			}
		case "leftDown":
			for i := range len(MAS) {
				r, c := v[0]-i, v[1]+i
				if matrix[r][c] != string(MAS[i]) {
					isMAS = false
				}
			}
		case "rightDown":
			for i := range len(MAS) {
				r, c := v[0]-i, v[1]-i
				if matrix[r][c] != string(MAS[i]) {
					isMAS = false
				}
			}

		}

		if isMAS {
			count++
		}
	}

	return count == 2
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
