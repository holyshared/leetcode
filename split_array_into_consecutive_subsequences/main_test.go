package main

import "testing"

func TestA(t *testing.T) {
	actual := isPossible([]int{1, 2, 3, 3, 4, 5})
	if !actual {
		t.Fatalf("actual = %v, expected = true", actual)
	}
}
