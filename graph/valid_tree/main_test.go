package main

import (
	"testing"
)

func TestA(t *testing.T) {
	actual := validTree(5, [][]int{
		{0, 1}, {0, 2}, {0, 3}, {1, 4},
	})

	if actual != true {
		t.Fatal("actual = false, expected = true")
	}
}

func TestB(t *testing.T) {
	actual := validTree(5, [][]int{
		{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4},
	})

	if actual != false {
		t.Fatal("actual = true, expected = false")
	}
}
