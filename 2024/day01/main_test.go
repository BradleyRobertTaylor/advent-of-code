package main

import "testing"

var exampleList1 = []int{3, 4, 2, 1, 3, 3}
var exampleList2 = []int{4, 3, 5, 3, 9, 3}

func TestPart1Example(t *testing.T) {
	expected := 11
	actual := Part1(exampleList1, exampleList2)

	if actual != expected {
		t.Errorf("Answer was incorrect. Got: %d, Expected: %d", actual, expected)
	}
}

func TestPart1(t *testing.T) {
}

// func TestPart2() {}
