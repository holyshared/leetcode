package main

import "testing"

func TestA(t *testing.T) {
	mat := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 4, 5, 8, 10},
		[]int{3, 5, 7, 9, 11},
		[]int{1, 3, 5, 7, 9},
	}

	actual := smallestCommonElement(mat)

	if actual != 5 {
		t.Fatalf("oops!!, %d", actual)
	}
}

func TestB(t *testing.T) {
	mat := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}

	actual := smallestCommonElement(mat)

	if actual != -1 {
		t.Fatalf("oops!!, %d", actual)
	}
}
