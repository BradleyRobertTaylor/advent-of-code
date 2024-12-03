package main

import (
	_ "embed"
	// "fmt"
	"regexp"
	"strings"

	u "github.com/BradleyRobertTaylor/advent-of-code/util"
)

//go:embed day3-input.txt
var input string

//go:embed day3-example.txt
var exampleInput string

//go:embed day3-example2.txt
var exampleInput2 string

func init() {
	input = strings.TrimRight(input, "\n")
	exampleInput = strings.TrimRight(exampleInput, "\n")
}

func Part1(input string) int {
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	pairs := u.Map(r.FindAllStringSubmatch(input, -1), func(matches []string) int {
		return u.ToInt(matches[1]) * u.ToInt(matches[2])
	})
	return u.Sum(pairs...)
}

func Part2(input string) int {
	r, _ := regexp.Compile(`don't\(\)`)
	dontIndices := r.FindAllIndex([]byte(input), -1)
	r, _ = regexp.Compile(`do\(\)`)
	doIndices := r.FindAllIndex([]byte(input), -1)

	var product int
	var temp string
	doPointer, dontPointer := 0, 0
	do := true
	for i, ch := range input {
		if do {
			temp += string(ch)
		}

		if dontPointer < len(dontIndices) && i == dontIndices[dontPointer][0] {
			product += Part1(temp)
			temp = ""
			dontPointer++
			do = false
		}

		if doPointer < len(doIndices) && i == doIndices[doPointer][0] {
			do = true
			doPointer++
		}
	}

	if do {
		product += Part1(temp)
	}

	return product
}
