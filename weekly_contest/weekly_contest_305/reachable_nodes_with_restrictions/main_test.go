package main

import "testing"

func TestA(t *testing.T) {
	actual := reachableNodes(7, [][]int{{0,1},{1,2},{3,1},{4,0},{0,5},{5,6}}, []int{4,5})
	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}

func TestB(t *testing.T) {
	actual := reachableNodes(7, [][]int{{0,1},{0,2},{0,5},{0,4},{3,2},{6,5}}, []int{4,2,1})
	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}

func TestC(t *testing.T) {
	actual := reachableNodes(2, [][]int{{0,1}}, []int{1})
	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}