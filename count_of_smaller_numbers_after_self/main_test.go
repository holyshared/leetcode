package main

import "testing"

func expect(a, e []int) bool {
	for i, v := range a {
		if v != e[i] {
			return false
		}
	}
	return true
}

func TestA(t *testing.T) {
	actual := countSmaller([]int{0,2,1})
	expected := []int{0,1,0}

	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}