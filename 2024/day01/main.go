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
	exampleInput = strings.TrimRight(input, "\n")
}

func main() {
	fmt.Println(Part1(input))
}

func Part1(input string) int {
	var list1, list2 []int
	listOfNumPairs := strings.Split(input, "\n")
	for _, l := range listOfNumPairs {
		nums := strings.Fields(l)

		list1 = append(list1, util.Num(nums[0]))
		list2 = append(list2, util.Num(nums[1]))
	}
	sort.Ints(list1)
	sort.Ints(list2)

	var result int
	for i := range len(list1) {
		distance := math.Abs(float64(list1[i]) - float64(list2[i]))
		result += int(distance)
	}

	return result
}

func Part2() {}
