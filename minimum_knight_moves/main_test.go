package main

import (
	"testing"
)

func Test21(t *testing.T) {
	actual := minKnightMoves(2, 1)
	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}

func Test55(t *testing.T) {
	actual := minKnightMoves(5, 5)
	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}
