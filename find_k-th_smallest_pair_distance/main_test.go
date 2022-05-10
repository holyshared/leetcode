package main

import (
	"testing"
)

func Test1112(t *testing.T) {
	actual := smallestDistancePair([]int{1, 1, 1}, 2)
	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}

func Test1613(t *testing.T) {
	actual := smallestDistancePair([]int{1, 6, 1}, 3)
	if actual != 5 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}
