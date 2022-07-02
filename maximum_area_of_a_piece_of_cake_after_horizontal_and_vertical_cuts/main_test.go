package main

import "testing"

func TestA(t *testing.T) {
	h := 5
	w := 4
	horizontalCuts := []int{1, 2, 4}
	verticalCuts := []int{1, 3}

	actual := maxArea(h, w, horizontalCuts, verticalCuts)

	if actual != 4 {
		t.Fatalf("acutal = %d, expected = 4", actual)
	}
}

func TestB(t *testing.T) {
	h := 5
	w := 4
	horizontalCuts := []int{3, 1}
	verticalCuts := []int{1}

	actual := maxArea(h, w, horizontalCuts, verticalCuts)

	if actual != 6 {
		t.Fatalf("acutal = %d, expected = 6", actual)
	}
}

func TestC(t *testing.T) {
	h := 5
	w := 4
	horizontalCuts := []int{3}
	verticalCuts := []int{3}

	actual := maxArea(h, w, horizontalCuts, verticalCuts)

	if actual != 9 {
		t.Fatalf("acutal = %d, expected = 9", actual)
	}
}

func TestD(t *testing.T) {
	h := 1000000000
	w := 1000000000
	horizontalCuts := []int{2}
	verticalCuts := []int{2}

	actual := maxArea(h, w, horizontalCuts, verticalCuts)

	if actual != 81 {
		t.Fatalf("acutal = %d, expected = 81", actual)
	}
}
