package main

import "testing"

func TestA(t *testing.T) {
	actual := minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7})
	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}

func TestB(t *testing.T) {
	actual := minSetSize([]int{7, 7, 7, 7, 7, 7})
	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}

func TestC(t *testing.T) {
	actual := minSetSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if actual != 5 {
		t.Fatalf("actual = %d, expected = 5", actual)
	}
}
