package main

import "testing"

func TestA(t *testing.T) {
	matrix := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	actual := shortestPathBinaryMatrix(matrix)

	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}

func TestB(t *testing.T) {
	matrix := [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	actual := shortestPathBinaryMatrix(matrix)

	if actual != -1 {
		t.Fatalf("actual = %d, expected = -1", actual)
	}
}

func TestC(t *testing.T) {
	matrix := [][]int{{0, 1}, {1, 0}}
	actual := shortestPathBinaryMatrix(matrix)

	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}
