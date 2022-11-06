package main

import "testing"

func TestA(t *testing.T) {
	actual := applyOperations([]int{1, 2, 2, 1, 1, 0})
	t.Fatalf("actual = %v expected = %v", actual, []int{1, 4, 2, 0, 0, 0})
}

func TestB(t *testing.T) {
	actual := applyOperations([]int{0, 1})
	t.Fatalf("actual = %v expected = %v", actual, []int{1, 0})
}
