package main

import "testing"

func TestA(t *testing.T) {
	answer := arithmeticTriplets([]int{0,1,4,6,7,10}, 3)
	if answer != 2 {
		t.Fatalf("actual = %d, expected = 2", answer)
	}
}