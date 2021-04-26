package main

import "testing"

func TestA(t *testing.T) {
	actual := furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1)

	if actual != 4 {
		t.Fatalf("index: %d != 4", actual)
	}
}

func TestB(t *testing.T) {
	actual := furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2)

	if actual != 7 {
		t.Fatalf("index: %d != 4", actual)
	}
}

func TestC(t *testing.T) {
	actual := furthestBuilding([]int{14, 3, 19, 3}, 17, 0)

	if actual != 3 {
		t.Fatalf("index: %d != 3", actual)
	}
}

func TestD(t *testing.T) {
	actual := furthestBuilding([]int{1, 2, 3, 4}, 0, 3)

	if actual != 3 {
		t.Fatalf("index: %d != 3", actual)
	}
}
