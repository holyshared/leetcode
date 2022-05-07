package main

import "testing"

func Test11(t *testing.T) {
	actual := intersection([]int{1}, []int{1})

	if len(actual) != 1 {
		t.Fatal("oops!!")
	}
}
