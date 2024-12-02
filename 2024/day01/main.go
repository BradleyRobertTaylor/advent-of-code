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
var list1, list2 []int

func init() {
	listOfNumPairs := strings.Split(input, "\n")
	for _, l := range listOfNumPairs {
		nums := strings.Fields(l)

		if len(nums) < 2 {
			continue
		}

		list1 = append(list1, util.Num(nums[0]))
		list2 = append(list2, util.Num(nums[1]))
	}
}

func main() {
	fmt.Println(Part1(list1, list2))
}

func Part1(list1, list2 []int) int {
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
