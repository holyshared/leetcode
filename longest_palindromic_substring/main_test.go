package main

import (
	"testing"
)

func TestValue_babad(t *testing.T) {
	actual := longestPalindrome("babad")

	if actual != "bab" {
		t.Fatalf("actual = %s, expected = bab", actual)
	}
}

func TestValue_cbbd(t *testing.T) {
	actual := longestPalindrome("cbbd")

	if actual != "bb" {
		t.Fatalf("actual = %s, expected = bb", actual)
	}
}
