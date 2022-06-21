package main

import (
	"testing"
)

func TestA(t *testing.T) {
	actual := earliestAcq([][]int{
		{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5},
	}, 6)

	if actual != 20190301 {
		t.Fatalf("acutal = %d, expected = 20190301", actual)
	}
}


func TestB(t *testing.T) {
	actual := earliestAcq([][]int{
		{0,2,0},{1,0,1},{3,0,3},{4,1,2},{7,3,1},
	}, 4)

	if actual != 3 {
		t.Fatalf("acutal = %d, expected = 3", actual)
	}
}
