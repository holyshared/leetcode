package main

import (
	"testing"
)

func TestNums111223K2(t *testing.T) {
	actual := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)

	if actual[0] != 1 {
		t.Fatalf("actual = %d, expected = 1", actual[0])
	}
	if actual[1] != 2 {
		t.Fatalf("actual = %d, expected = 2", actual[1])
	}
}

func TestNums1K1(t *testing.T) {
	actual := topKFrequent([]int{1}, 1)

	if actual[0] != 1 {
		t.Fatalf("actual = %d, expected = 1", actual[0])
	}
}

func TestNums1112223345555K3(t *testing.T) {
	actual := topKFrequent([]int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 5, 5, 5}, 3)
	if actual[0] != 5 {
		t.Fatalf("actual = %d, expected = 5", actual[0])
	}
	if actual[1] != 1 {
		t.Fatalf("actual = %d, expected = 1", actual[1])
	}
	if actual[2] != 2 {
		t.Fatalf("actual = %d, expected = 2", actual[2])
	}
}
