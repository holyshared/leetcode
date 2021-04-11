package main

import "testing"

func TestA(t *testing.T) {
	source := []byte{'h', 'e', 'l', 'l', 'o'}
	expected := []byte{'o', 'l', 'l', 'e', 'h'}
	reverseString(source)

	for i, v := range source {
		if v != expected[i] {
			t.Fatalf("oops!! %d", i)
		}
	}
}

func TestB(t *testing.T) {
	source := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	expected := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	reverseString(source)

	for i, v := range source {
		if v != expected[i] {
			t.Fatalf("oops!! %d", i)
		}
	}
}
