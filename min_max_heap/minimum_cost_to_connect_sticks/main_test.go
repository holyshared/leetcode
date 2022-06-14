package main

import (
	"testing"
)

func TestNums243(t *testing.T) {
	actual := connectSticks([]int{2, 4, 3})

	if actual != 14 {
		t.Fatalf("actual = %d, expected = 14", actual)
	}
}

func TestNums1835(t *testing.T) {
	actual := connectSticks([]int{1, 8, 3, 5})

	if actual != 30 {
		t.Fatalf("actual = %d, expected = 30", actual)
	}
}

func TestNums5(t *testing.T) {
	actual := connectSticks([]int{5})

	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}
