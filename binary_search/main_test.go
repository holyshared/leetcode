package main

import "testing"

func TestA(t *testing.T) {
	actual := search([]int{-1, 0, 3, 5, 9, 12}, 9)

	if actual != 4 {
		t.Fatalf("oops!! %d", actual)
	}

}

func TestB(t *testing.T) {
	actual := search([]int{-1, 0, 3, 5, 9, 12}, 2)

	if actual != -1 {
		t.Fatalf("oops!! %d", actual)
	}
}
