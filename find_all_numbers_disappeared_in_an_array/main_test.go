package main

import "testing"

func Test11(t *testing.T) {
	expeted := []int{2}
	actual := findDisappearedNumbers([]int{1, 1})

	for i := 0; i < len(expeted); i++ {
		if actual[i] != expeted[i] {
			t.Fatalf("%d", actual)
		}
	}
}
