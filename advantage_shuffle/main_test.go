package main

import "testing"

func TestA(t *testing.T) {
	actual := advantageCount([]int{2,7,11,15}, []int{1,10,4,11})
	expext := []int{2,11,7,15}

	for i, v := range actual {
		if v != expext[i] {
			t.Fatalf("%d %d", v, expext[i])
		}
	}
}