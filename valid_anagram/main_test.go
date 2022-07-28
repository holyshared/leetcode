package main

import "testing"

func TestA(t *testing.T) {
	if !isAnagram("az", "az") {
		t.Fatalf("expected = true")
	}
}

func TestB(t *testing.T) {
	if !isAnagram("anagram", "nagaram") {
		t.Fatalf("expected = true")
	}
}
