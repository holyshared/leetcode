package main

import "testing"

func TestA(t *testing.T) {
	actual := findPaths(2, 2, 2, 0, 0)

	if actual != 6 {
		t.Fatalf("actual = %d, expected = 6", actual)
	}
}

func TestB(t *testing.T) {
	actual := findPaths(1, 3, 3, 0, 1)

	if actual != 12 {
		t.Fatalf("actual = %d, expected = 12", actual)
	}
}


func TestC(t *testing.T) {
	actual := findPaths(10, 10, 3, 0, 0)

	if actual != 10 {
		t.Fatalf("actual = %d, expected = 10", actual)
	}
}

func TestD(t *testing.T) {
	actual := findPaths(8, 50, 23, 5, 26)

	if actual != 914783380 {
		t.Fatalf("actual = %d, expected = 914783380", actual)
	}
}

func TestE(t *testing.T) {
	actual := findPaths(36, 5, 50, 15, 3)

	if actual != 390153306 {
		t.Fatalf("actual = %d, expected = 390153306", actual)
	}
}
