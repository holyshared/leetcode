package main

import (
	"testing"
)

func TestNums13K8(t *testing.T) {
	actual := kthSmallest([][]int{
		[]int{1, 5, 9},
		[]int{10, 11, 13},
		[]int{12, 13, 15},
	}, 8)

	if actual != 13 {
		t.Fatalf("actual = %d, expected = 13", actual)
	}
}

func TestNumsN1K1(t *testing.T) {
	actual := kthSmallest([][]int{
		[]int{-5},
	}, 1)

	if actual != -5 {
		t.Fatalf("actual = %d, expected = -5", actual)
	}
}
