package main

import "testing"

func TestA(t * testing.T) {
	matrix := [][]int{{1,5,9},{10,11,13},{12,13,15}}
	actual := kthSmallest(matrix, 8)

	if actual != 13 {
		t.Fatalf("actual = %d, expected = 13", actual)
	}
}

func TestB(t * testing.T) {
	matrix := [][]int{{-5}}
	actual := kthSmallest(matrix, 1)

	if actual != -5 {
		t.Fatalf("actual = %d, expected = -5", actual)
	}
}