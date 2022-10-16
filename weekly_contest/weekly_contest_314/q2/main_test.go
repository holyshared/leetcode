package main

import "testing"

func TestA(t *testing.T) {
	actual := findArray([]int{5, 2, 0, 3, 1})
	t.Fatalf("actual = %v", actual)
}
