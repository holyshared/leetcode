package main

import "testing"

func TestA(t *testing.T) {
	colors := []int{1, 1, 2, 1, 3, 2, 2, 3, 3}
	queries := [][]int{{1, 3}, {2, 2}, {6, 1}}

	actual := shortestDistanceColor(colors, queries)

	if actual[0] != 3 {
		t.Fatalf("actual = %d, expected = 3", actual[0])
	}
	if actual[1] != 0 {
		t.Fatalf("actual = %d, expected = 0", actual[1])
	}
	if actual[2] != 3 {
		t.Fatalf("actual = %d, expected = 3", actual[2])
	}
}

func TestB(t *testing.T) {
	colors := []int{1, 2}
	queries := [][]int{{0, 3}}

	actual := shortestDistanceColor(colors, queries)

	if actual[0] != -1 {
		t.Fatalf("actual = %d, expected = -1", actual[0])
	}
}

func TestC(t *testing.T) {
	colors := []int{1, 1, 2, 1, 3, 2, 2, 3, 3}
	queries := [][]int{{2, 1}}

	actual := shortestDistanceColor(colors, queries)

	if actual[0] != 1 {
		t.Fatalf("actual = %d, expected = 1", actual[0])
	}
}
