package main

import (
	"testing"
)

func TestNums321564K2(t *testing.T) {
	actual := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)

	if actual != 5 {
		t.Fatalf("actual = %d, expected = 5", actual)
	}
}

func TestNums323124556K4(t *testing.T) {
	actual := findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)

	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}
