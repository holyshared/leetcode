package main

import "testing"

func TestA(t *testing.T) {
	actual := largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1})
	if actual != 8 {
		t.Fatalf("8 != %d", actual)
	}
}

func TestB(t *testing.T) {
	actual := largestUniqueNumber([]int{9, 9, 8, 8})
	if actual != -1 {
		t.Fatalf("-1 != %d", actual)
	}
}

func TestC(t *testing.T) {
	actual := largestUniqueNumber([]int{9, 9, 8})
	if actual != 8 {
		t.Fatalf("8 != %d", actual)
	}
}
