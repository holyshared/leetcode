package main

import "testing"

func TestA(t *testing.T) {
	actual := countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}})
	if actual != 2 {
		t.Fatalf("actual: %d, expcted: 2", actual)
	}
}

func TestB(t *testing.T) {
	actual := countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}})
	if actual != 1 {
		t.Fatalf("actual: %d, expcted: 1", actual)
	}
}
