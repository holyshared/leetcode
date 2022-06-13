package main

import (
	"testing"
)

func TestK122(t *testing.T) {
	actual := kClosest([][]int{
		{1, 3},
		{-2, 2},
	}, 1)

	if actual[0][0] != -2 {
		t.Fatalf("actual[0][0] = %d, expected = -2", actual[0][0])
	}

	if actual[0][1] != 2 {
		t.Fatalf("actual[0][1] = %d, expected = 2", actual[0][1])
	}
}

func TestK233m24(t *testing.T) {
	actual := kClosest([][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}, 2)

	if actual[0][0] != 3 {
		t.Fatalf("actual[0][0] = %d, expected = 3", actual[0][0])
	}

	if actual[0][1] != 3 {
		t.Fatalf("actual[0][1] = %d, expected = 3", actual[0][1])
	}

	if actual[1][0] != -2 {
		t.Fatalf("actual[1][0] = %d, expected = -2", actual[1][0])
	}

	if actual[1][1] != 4 {
		t.Fatalf("actual[1][1] = %d, expected = 4", actual[1][1])
	}
}
