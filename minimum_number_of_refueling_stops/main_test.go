package main

import "testing"

func TestA(t *testing.T) {
	actual := minRefuelStops(1, 1, [][]int{})
	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}

func TestB(t *testing.T) {
	actual := minRefuelStops(100, 1, [][]int{{10, 100}})
	if actual != -1 {
		t.Fatalf("actual = %d, expected = -1", actual)
	}
}

func TestC(t *testing.T) {
	actual := minRefuelStops(100, 10, [][]int{
		{10, 60}, {20, 30}, {30, 30}, {60, 40},
	})
	// [10 70 110 140 170]
	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}
