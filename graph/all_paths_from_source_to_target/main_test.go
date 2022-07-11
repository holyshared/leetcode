package main

import "testing"

func TestA(t *testing.T) {
	actual := allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}})

	if len(actual) != 2 {
		t.Fatalf("actual routes = %d", len(actual))
	}
}

func TestB(t *testing.T) {
	actual := allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}})

	if len(actual) != 5 {
		t.Fatalf("actual routes = %d", len(actual))
	}
}

func TestC(t *testing.T) {
	actual := allPathsSourceTarget([][]int{{4,3,1},{3,2,4},{},{4},{}})

	if len(actual) != 4 {
		t.Fatalf("actual routes = %d", len(actual))
	}
}

/**
 0 -> 2 -> 1
 */
func TestD(t *testing.T) {
	actual := allPathsSourceTarget([][]int{{2},{},{1}})

	if len(actual) != 1 {
		t.Fatalf("actual routes = %d", len(actual))
	}
}
