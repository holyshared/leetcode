package main

import "testing"

func TestA(t *testing.T) {
	actual := numSubmatrixSumTarget([][]int{
		{0, 1, 0}, {1, 1, 1}, {0, 1, 0},
	}, 0)

	if actual != 4 {
		t.Fatalf("%d != 4", actual)
	}
}

func TestB(t *testing.T) {
	actual := numSubmatrixSumTarget([][]int{
		{1, -1}, {-1, 1},
	}, 0)

	if actual != 5 {
		t.Fatalf("%d != 5", actual)
	}
}
