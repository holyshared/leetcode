package main

import "testing"

func Test24(t *testing.T) {
	if numFactoredBinaryTrees([]int{2, 4}) != 3 {
		t.Fatal("oops!!")
	}
}

func Test24510(t *testing.T) {
	actual := numFactoredBinaryTrees([]int{2, 4, 5, 10})
	if actual != 7 {
		t.Fatalf("oops!! %d", actual)
	}
}

func Test18362(t *testing.T) {
	actual := numFactoredBinaryTrees([]int{18, 3, 6, 2})
	if actual != 12 {
		t.Fatalf("oops!! %d", actual)
	}
}

func TestWrong(t *testing.T) {
	actual := numFactoredBinaryTrees([]int{2, 82, 64, 18, 85, 86, 81, 4, 67, 95})
	if actual != 11 {
		t.Fatalf("oops!! %d", actual)
	}
}
