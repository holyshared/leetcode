package main

import "testing"

func TestA(t *testing.T) {
	actual := uniquePaths(3, 2)

	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}

func TestB(t *testing.T) {
	actual := uniquePaths(3, 7)

	if actual != 28 {
		t.Fatalf("actual = %d, expected = 28", actual)
	}
}

func TestC(t *testing.T) {
	actual := uniquePaths(23, 12)

	if actual != 193536720 {
		t.Fatalf("actual = %d, expected = 193536720", actual)
	}
}