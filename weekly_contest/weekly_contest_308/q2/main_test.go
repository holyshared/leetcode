package main

import "testing"

func TestA(t *testing.T) {
	acutual := removeStars("leet**cod*e")

	if acutual != "lecoe" {
		t.Fatalf("actual = %s, expected = lecoe", acutual)
	}

}
