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
		{"Day 2 Example input", exampleInput, 2},
		{"Day 2 Actual input", input, 483},
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
		{"Day 2 Example input", exampleInput, 4},
		{"Day 2 Actual input", input, 528},
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
