package main

import (
	"testing"
)

func TestA(t *testing.T) {
	expected := [][]int{
		[]int{0, 4},
		[]int{1, 3},
		[]int{1, 4},
		[]int{2, 2},
		[]int{3, 0},
		[]int{3, 1},
		[]int{4, 0},
	}

	actual := pacificAtlantic([][]int{
		[]int{1, 2, 2, 3, 5},
		[]int{3, 2, 3, 4, 4},
		[]int{2, 4, 5, 3, 1},
		[]int{6, 7, 1, 4, 5},
		[]int{5, 1, 1, 2, 4},
	})

	for i := 0; i < len(actual); i++ {
		if actual[i][0] != expected[i][0] || actual[i][1] != expected[i][1] {
			t.Fatalf("(%d, %d) != (%d, %d)", actual[i][0], actual[i][1], expected[i][0], expected[i][1])
		}
	}
}
