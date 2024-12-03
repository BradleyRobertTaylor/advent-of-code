package main

import (
	_ "embed"
	"math"
	"strings"

	u "github.com/BradleyRobertTaylor/advent-of-code/util"
)

//go:embed day2-input.txt
var input string

//go:embed day2-example.txt
var exampleInput string

func init() {
	input = strings.TrimRight(input, "\n")
	exampleInput = strings.TrimRight(exampleInput, "\n")
}

func Part1(input string) int {
	reports := getReports(input)
	var safe int

	for _, r := range reports {
		if isReportSafe(r) {
			safe++
		}
	}

	return safe
}

// Not very efficient n^2 :(
func Part2(input string) int {
	reports := getReports(input)
	var safe int

	for _, r := range reports {
		if isReportSafe(r) {
			safe++
			continue
		}

		for i := range r {
			sliceWithRemoved := make([]int, 0, len(r)-1)
			for j := range r {
				if j != i {
					sliceWithRemoved = append(sliceWithRemoved, r[j])
				}
			}

			if isReportSafe(sliceWithRemoved) {
				safe++
				break
			}
		}
	}

	return safe
}

func getReports(input string) [][]int {
	inputSlice := strings.Split(input, "\n")
	reports := u.Map(inputSlice, func(s string) []int {
		return u.Map(strings.Fields(s), func(s string) int {
			return u.ToInt(s)
		})
	})

	return reports
}

func isReportSafe(r []int) bool {
	isInc := r[0] < r[1]

	for i := 1; i < len(r); i++ {
		n := r[i]
		diff := r[i-1] - n
		absDiff := math.Abs(float64(diff))

		if diff == 0 || absDiff > 3 {
			return false
		}

		if diff < 0 && isInc == false {
			return false
		}

		if diff > 0 && isInc == true {
			return false
		}
	}

	return true
}
