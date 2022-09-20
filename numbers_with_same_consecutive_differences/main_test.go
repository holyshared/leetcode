package main

import (
	"sort"
	"testing"
)

func expectSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	compA := func(i, j int) bool {
		return a[i] < a[j]
	}
	compB := func(i, j int) bool {
		return b[i] < b[j]
	}
	sort.Slice(a, compA)
	sort.Slice(b, compB)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestA(t *testing.T) {
	actual := numsSameConsecDiff(3, 7)
	expected := []int{181, 292, 707, 818, 929}

	if !expectSlice(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}

func TestB(t *testing.T) {
	actual := numsSameConsecDiff(2, 1)
	expected := []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}

	if !expectSlice(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}

func TestC(t *testing.T) {
	actual := numsSameConsecDiff(2, 0)
	expected := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}

	if !expectSlice(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}
