package main

import (
	_ "embed"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name, input string
		expected    int
	}{
		{"Day 5 Example input", exampleInput, 143},
		{"Day 5 Actual input", input, 4281},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)
			fmt.Println(tt.input)

			if got != tt.expected {
				t.Errorf("Answer was incorrect. Got: %d, Expected: %d", got, tt.expected)
			}
		})
	}
}

// func TestPart2(t *testing.T) {
// 	tests := []struct {
// 		name, input string
// 		expected    int
// 	}{
// 		{"Day 5 Example input", exampleInput, 9},
// 		{"Day 5 Actual input", input, 1990},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := Part2(tt.input)
//
// 			if got != tt.expected {
// 				t.Errorf("Answer was incorrect. Got: %d, Expected: %d", got, tt.expected)
// 			}
// 		})
// 	}
// }
