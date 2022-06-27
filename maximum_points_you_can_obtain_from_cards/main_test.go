package main

import (
	"testing"
)

func TestCP1234561_K3(t *testing.T) {
	actual := maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3)
	if actual != 12 {
		t.Fatalf("actual = %d, expected = 12", actual)
	}
}
