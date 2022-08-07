package main

import "testing"

func TestA(t *testing.T) {
	actual := networkDelayTime([][]int{{1, 2, 1}}, 2, 2)

	if actual != -1 {
		t.Fatalf("actual = %d, expected = -1", actual)
	}
}
