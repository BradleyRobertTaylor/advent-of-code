package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/BradleyRobertTaylor/advent-of-code/util"
)

//go:embed day1-input.txt
var input string

//go:embed day1-example.txt
var exampleInput string

func init() {
	input = strings.TrimRight(input, "\n")
	exampleInput = strings.TrimRight(exampleInput, "\n")
}

func main() {
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input string) int {
	var result float64
	leftList, rightList := getLists(input)
	for i := range len(leftList) {
		distance := math.Abs(leftList[i] - rightList[i])
		result += distance
	}

	return int(result)
}

func Part2(input string) int {
	leftList, rightList := getLists(input)

	count := make(map[float64]int)
	for _, n := range rightList {
		count[n]++
	}

	var result int
	for _, n := range leftList {
		if numCount, ok := count[n]; ok {
			result += numCount * int(n)
		}
	}

	return result
}

func getLists(input string) ([]float64, []float64) {
	var leftList, rightList []float64
	listOfNumPairs := strings.Split(input, "\n")

	for _, l := range listOfNumPairs {
		nums := strings.Fields(l)

		leftList = append(leftList, util.ToFloat64(nums[0]))
		rightList = append(rightList, util.ToFloat64(nums[1]))
	}
	sort.Float64s(leftList)
	sort.Float64s(rightList)

	return leftList, rightList
}
