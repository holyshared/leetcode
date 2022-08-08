package main

import "testing"

func TestA(t *testing.T) {
	actual := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}
