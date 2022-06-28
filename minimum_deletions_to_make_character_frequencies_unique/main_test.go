package main

import "testing"

func TestSaab(t *testing.T) {
	actual := minDeletions("aab")

	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}

func TestSaaabbbcc(t *testing.T) {
	actual := minDeletions("aaabbbcc")

	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}

func TestSceabaacb(t *testing.T) {
	actual := minDeletions("ceabaacb")

	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}

