package main

import "testing"

func expect(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestA(t *testing.T) {
	expected := []int{2, 3, 4}
	acutual := answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21})

	if !expect(acutual, expected) {
		t.Fatalf("actual = %v, expected = %v", acutual, expected)
	}

}
