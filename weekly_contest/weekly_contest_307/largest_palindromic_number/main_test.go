package main

import "testing"

func TestA(t *testing.T) {
	actual := largestPalindromic("444947137")
	if actual != "7449447" {
		t.Fatalf("actual = %s, expected = 7449447", actual)
	}

}

func TestB(t *testing.T) {
	actual := largestPalindromic("00009")
	if actual != "9" {
		t.Fatalf("actual = %s, expected = 9", actual)
	}

}
