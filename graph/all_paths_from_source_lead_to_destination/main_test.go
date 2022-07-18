package main

import "testing"

func TestA(t *testing.T) {
	actual := leadsToDestination(3, [][]int{{0, 1}, {0, 2}}, 0, 2)
	if actual != false {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}

func TestB(t *testing.T) {
	actual := leadsToDestination(4, [][]int{{0, 1}, {0, 3}, {1, 2}, {2, 1}}, 0, 3)
	if actual != false {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}

func TestC(t *testing.T) {
	actual := leadsToDestination(4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, 0, 3)
	if actual != true {
		t.Fatalf("actual = %v, expected = true", actual)
	}
}
