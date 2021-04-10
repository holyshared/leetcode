package main

import "testing"

func TestA(t *testing.T) {
	actual := longestIncreasingPath([][]int{
		[]int{9, 9, 4},
		[]int{6, 6, 8},
		[]int{2, 1, 1},
	})

	if actual != 4 {
		t.Fatal("oops!!")
	}
}
