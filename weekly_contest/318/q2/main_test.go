package main

import "testing"

func TestA(t *testing.T) {
	actual := maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3)
	if actual != 15 {
		t.Fatalf("actual = %v expected = %v", actual, 15)
	}
}

func TestB(t *testing.T) {
	actual := maximumSubarraySum([]int{4, 4, 4}, 3)
	if actual != 0 {
		t.Fatalf("actual = %v expected = %v", actual, 0)
	}
}

func TestC(t *testing.T) {
	actual := maximumSubarraySum([]int{-1, -2, -3, 1, 2, 3}, 3)
	if actual != 6 {
		t.Fatalf("actual = %v expected = %v", actual, 6)
	}
}

func TestD(t *testing.T) {
	actual := maximumSubarraySum([]int{9, 9, 9, 1, 2, 3}, 3)
	if actual != 12 {
		t.Fatalf("actual = %v expected = %v", actual, 12)
	}
}
