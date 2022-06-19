package main

import "testing"

func TestEdges5_011234(t *testing.T) {
	actual := countComponents(5, [][]int{
		{0, 1}, {1, 2}, {3, 4},
	})
	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}

func TestEdges5_01122334(t *testing.T) {
	actual := countComponents(5, [][]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 4},
	})
	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}
