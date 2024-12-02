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
}

func Part1(input string) int {
	var list1, list2 []float64
	listOfNumPairs := strings.Split(input, "\n")

	for _, l := range listOfNumPairs {
		nums := strings.Fields(l)

		list1 = append(list1, util.ToFloat64(nums[0]))
		list2 = append(list2, util.ToFloat64(nums[1]))
	}
	sort.Float64s(list1)
	sort.Float64s(list2)

	var result float64
	for i := range len(list1) {
		distance := math.Abs(list1[i] - list2[i])
		result += distance
	}

	return int(result)
}

func Part2() {}
