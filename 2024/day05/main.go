package main

import (
	_ "embed"
	"strings"

	"github.com/BradleyRobertTaylor/advent-of-code/util"
)

//go:embed day5-input.txt
var input string

//go:embed day5-example.txt
var exampleInput string

func init() {
	input = strings.TrimRight(input, "\n")
	exampleInput = strings.TrimRight(exampleInput, "\n")
}

func Part1(input string) int {
	r, updates := parseInput(input)
	rules := make(map[int][]int)

	for _, rule := range r {
		k, v := rule[0], rule[1]
		rules[k] = append(rules[k], v)
	}

	var result int

outer:
	for _, u := range updates {
		seen := make(map[int]bool)
		for _, num := range u {
			for _, notAllowedNum := range rules[num] {
				if _, ok := seen[notAllowedNum]; ok {
					continue outer
				}
			}

			seen[num] = true
		}

		result += u[len(u)/2]
	}

	return result
}

func Part2(input string) int {
	return 0
}

func parseInput(input string) ([][2]int, [][]int) {
	var (
		rules   [][2]int
		updates [][]int
	)

	sections := strings.Split(input, "\n\n")
	r, u := strings.Split(sections[0], "\n"), strings.Split(sections[1], "\n")
	for _, s := range u {
		nums := strings.Split(s, ",")
		level := make([]int, 0)

		for _, n := range nums {
			level = append(level, util.ToInt(n))
		}

		updates = append(updates, level)
	}

	for _, p := range r {
		pair := strings.Split(p, "|")
		rules = append(rules, [2]int{util.ToInt(pair[0]), util.ToInt(pair[1])})
	}

	return rules, updates
}
