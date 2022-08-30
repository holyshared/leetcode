package main

import "testing"

// 4 + (6 + 3)      10
// 6 1 3 1 -> 11
func TestA(t *testing.T) {
	acutual := garbageCollection(
		[]string{"G", "P", "GP", "GG"},
		[]int{2, 4, 3},
	)

	if acutual != 21 {
		t.Fatalf("actual = %d, expected = 21", acutual)
	}

}
