package main

import "testing"

func TestA(t *testing.T) {
	actual := guessNumber(10)
	if actual != 6 {
		t.Fatalf("%d", actual)
	}
}
