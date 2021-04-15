package main

import "testing"

func TestA(t *testing.T) {
	actual := minSwaps([]int{1, 0, 1, 0, 1})

	if actual != 1 {
		t.Fatalf("%d != 1", actual)
	}

}
