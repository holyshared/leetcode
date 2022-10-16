package main

import "testing"

func TestA(t *testing.T) {
	actual := hardestWorker(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}})
	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}
func TestB(t *testing.T) {
	actual := hardestWorker(26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}})
	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}

func TestC(t *testing.T) {
	actual := hardestWorker(2, [][]int{{0, 10}, {1, 20}})
	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}
