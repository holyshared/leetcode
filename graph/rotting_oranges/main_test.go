package main

import "testing"

func TestA(t *testing.T) {
	actual := orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})

	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}

func TestB(t *testing.T) {
	actual := orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}})

	if actual != -1 {
		t.Fatalf("actual = %d, expected = -1", actual)
	}
}

func TestC(t *testing.T) {
	actual := orangesRotting([][]int{{0, 2}})

	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}
