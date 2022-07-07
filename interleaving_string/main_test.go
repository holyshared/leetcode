package main

import "testing"

func TestA(t *testing.T) {
	actual := isInterleave("aabcc", "dbbca", "aadbbcbcac")

	if actual != true {
		t.Fatalf("actual = false")
	}
}
