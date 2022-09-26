package main

import "testing"

func TestA(t *testing.T) {
	actual := findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})
	if actual != 3 {
		t.Fatalf("actual = %v, expected = 3", actual)
	}
}
