package main

import (
	"testing"
)

func TestNums274181(t *testing.T) {
	actual := lastStoneWeight([]int{2,7,4,1,8,1})

	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}
