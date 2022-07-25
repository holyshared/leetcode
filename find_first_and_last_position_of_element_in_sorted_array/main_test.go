package main

import "testing"

func TestA(t *testing.T) {
	actual := searchRange([]int{5,7,7,8,8,10}, 8)

	if actual[0] != 3 {
		t.Fatalf("actual = %d, expected = 3", actual[0])
	}

	if actual[1] != 4 {
		t.Fatalf("actual = %d, expected = 4", actual[1])
	}
}

func TestB(t *testing.T) {
	actual := searchRange([]int{5,7,7,8,8,10}, 6)

	if actual[0] != -1 {
		t.Fatalf("actual = %d, expected = 3", actual[0])
	}

	if actual[1] != -1 {
		t.Fatalf("actual = %d, expected = 4", actual[1])
	}
}

func TestC(t *testing.T) {
	actual := searchRange([]int{}, 0)

	if actual[0] != -1 {
		t.Fatalf("actual = %d, expected = 3", actual[0])
	}

	if actual[1] != -1 {
		t.Fatalf("actual = %d, expected = 4", actual[1])
	}
}
