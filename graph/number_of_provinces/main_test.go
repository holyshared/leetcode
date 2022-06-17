package main

import (
	"testing"
)

func TestA(t *testing.T) {
	actual := findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	})

	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}

func TestB(t *testing.T) {
	actual := findCircleNum([][]int{
		{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
	})

	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}
