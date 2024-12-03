package main

import (
	_ "embed"
	"regexp"
	"strings"

	"github.com/BradleyRobertTaylor/advent-of-code/util"
)

//go:embed day3-input.txt
var input string

//go:embed day3-example.txt
var exampleInput string

func init() {
	input = strings.TrimRight(input, "\n")
	exampleInput = strings.TrimRight(exampleInput, "\n")
}

func Part1(input string) int {
	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllString(input, -1)
	l, _ := regexp.Compile(`\d{1,3},`)
	r, _ = regexp.Compile(`,\d{1,3}`)

	var product int
	for _, m := range matches {
		left, right := l.FindString(m), r.FindString(m)
		left = left[:len(left)-1]
		right = right[1:]

		product += util.ToInt(left) * util.ToInt(right)
	}

	return product
}

func Part2(input string) int {
	return 0
}
