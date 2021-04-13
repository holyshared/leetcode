package main

import "testing"

func TestA(t *testing.T) {
	actual := mySqrt(8)
	if actual != 2 {
		t.Fatalf("%d", actual)
	}
}
