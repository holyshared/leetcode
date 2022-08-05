package main

import (
	"testing"
)

func TestA(t *testing.T) {
	actual := combinationSum4([]int{1, 2, 3}, 4)

	if actual != 7 {
		t.Fatalf("actual = %d, expected = 7", actual)
	}
}

func TestB(t *testing.T) {
	actual := combinationSum4([]int{9}, 3)

	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}
