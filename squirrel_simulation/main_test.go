package main

import "testing"

func TestA(t *testing.T) {
	distance := minDistance(
		5,
		7,
		[]int{2, 2},
		[]int{4, 4},
		[][]int{{3, 0}, {2, 5}},
	)
	if distance != 12 {
		t.Fatalf("%d", distance)
	}
}

func TestB(t *testing.T) {
	distance := minDistance(
		5,
		7,
		[]int{2, 2},
		[]int{4, 4},
		[][]int{{3, 0}, {2, 5}, {0, 0}},
	)
	if distance != 20 {
		t.Fatalf("%d", distance)
	}
}
