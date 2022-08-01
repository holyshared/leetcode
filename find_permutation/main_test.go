package main

import "testing"

func TestA(t *testing.T) {
	actual := findPermutation("I")

	if actual[0] != 1 {
		t.Fatalf("actual = %v", actual)
	}
	if actual[1] != 2 {
		t.Fatalf("actual = %v", actual)
	}


}

func TestB(t *testing.T) {
	actual := findPermutation("DI")
	if actual[0] != 2 {
		t.Fatalf("actual = %v", actual)
	}
	if actual[1] != 1 {
		t.Fatalf("actual = %v", actual)
	}
	if actual[2] != 3 {
		t.Fatalf("actual = %v", actual)
	}
}