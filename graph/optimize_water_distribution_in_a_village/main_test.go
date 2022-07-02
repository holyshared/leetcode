package main

import (
	"testing"
)

func TestN3Wells122(t *testing.T) {
	n := 3
	wells := []int{1, 2, 2}
	pipes := [][]int{{1, 2, 1}, {2, 3, 1}}

	actual := minCostToSupplyWater(n, wells, pipes)

	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}

func TestN2Wells11(t *testing.T) {
	n := 2
	wells := []int{1, 1}
	pipes := [][]int{{1, 2, 1}, {1, 2, 2}}

	actual := minCostToSupplyWater(n, wells, pipes)

	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}
