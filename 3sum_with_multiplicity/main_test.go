package main

import "testing"

func TestA(t *testing.T) {
	acutual := threeSumMulti([]int{2, 2, 3, 2}, 7)
	if acutual != 3 {
		t.Fatalf("%d", acutual)
	}
}
