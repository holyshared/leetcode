package main

import (
	"testing"
)

func Test13_2(t *testing.T) {
	actual := findMedianSortedArrays([]int{1, 3}, []int{2})
	if actual != 2.0 {
		t.Fatal("oops!!")
	}
}

func Test12_34(t *testing.T) {
	actual := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	if actual != 2.5 {
		t.Fatal("oops!!")
	}
}
