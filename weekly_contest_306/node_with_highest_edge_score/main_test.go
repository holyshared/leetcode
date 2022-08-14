package main

import "testing"

func TestA(t *testing.T) {
	actual := edgeScore([]int{1,0,0,0,0,7,7,5})

	if actual != 7 {
		t.Fatalf("actual = %d, expected = 7", actual)
	}
}

func TestB(t *testing.T) {
	actual := edgeScore([]int{2,0,0,2})






	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}