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
		{"Example", exampleInput, 11},
		{"Actual", input, 2164381},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)
			t.Log(tt.input)

			if got != tt.expected {
				t.Errorf("Answer was incorrect. Got: %d, Expected: %d", got, tt.expected)
			}
		})
	}
}

// func TestPart2() {}
