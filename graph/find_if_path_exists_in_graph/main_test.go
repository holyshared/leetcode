package main

import "testing"

func TestA(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	source := 0
	destination := 2

	actual := validPath(n, edges, source, destination)
	if actual != true {
		t.Fatalf("actual = %v, expected = true", actual)
	}
}

func TestB(t *testing.T) {
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	source := 0
	destination := 5

	actual := validPath(n, edges, source, destination)
	if actual != false {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}
