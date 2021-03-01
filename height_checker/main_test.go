package main

import "testing"

func Test114213(t *testing.T) {
	changedCount := heightChecker([]int{1, 1, 4, 2, 1, 3})

	if changedCount != 3 {
		t.Fatalf("failed: %d", changedCount)
	}
}
