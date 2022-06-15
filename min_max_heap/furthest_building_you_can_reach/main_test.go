package main

import (
	"testing"
)

func Test427691412B5R1(t *testing.T) {
	actual := furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1)

	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}

func Test4122731820319B10R2(t *testing.T) {
	actual := furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2)

	if actual != 7 {
		t.Fatalf("actual = %d, expected = 7", actual)
	}
}

func Test143193B17R0(t *testing.T) {
	actual := furthestBuilding([]int{14, 3, 19, 3}, 17, 0)

	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}
