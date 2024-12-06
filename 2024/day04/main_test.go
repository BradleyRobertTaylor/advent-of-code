package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name, input string
		expected    int
	}{
		{"Day 4 Example input", exampleInput, 18},
		{"Day 4 Actual input", input, 2654},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)

			if got != tt.expected {
				t.Errorf("Answer was incorrect. Got: %d, Expected: %d", got, tt.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name, input string
		expected    int
	}{
		{"Day 4 Example input", exampleInput, 9},
		// {"Day 4 Actual input", input, 93465710},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)

			if got != tt.expected {
				t.Errorf("Answer was incorrect. Got: %d, Expected: %d", got, tt.expected)
			}
		})
	}
}
