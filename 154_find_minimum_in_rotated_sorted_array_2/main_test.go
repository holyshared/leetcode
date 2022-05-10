package main

import "testing"

func Test22201(t *testing.T) {
	actual := findMin([]int{2, 2, 2, 0, 1})
	if actual != 0 {
		t.Fatal("oops!!")
	}
}
