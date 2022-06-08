package main

import (
	"testing"
)

func TestMat20314K3(t *testing.T) {
	actual := kWeakestRows([][]int{
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 1, 1, 0},
		[]int{1, 0, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 1, 1, 1},
	}, 3)

	if actual[0] != 2 {
		t.Fatalf("actual[0] = %d, expected = 2", actual[0])
	}

	if actual[1] != 0 {
		t.Fatalf("actual[1] = %d, expected = 0", actual[1])
	}

	if actual[2] != 3 {
		t.Fatalf("actual[2] = %d, expected = 3", actual[2])
	}

}
