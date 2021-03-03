package main

import "testing"

func Test301(t *testing.T) {
	actual := missingNumber([]int{3, 0, 1})
	if actual != 2 {
		t.Fatalf("%d", actual)
	}
}

func Test01(t *testing.T) {
	actual := missingNumber([]int{0, 1})
	if actual != 2 {
		t.Fatalf("%d", actual)
	}
}

func Test964235701(t *testing.T) {
	actual := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	if actual != 8 {
		t.Fatalf("%d", actual)
	}
}

func Test0(t *testing.T) {
	actual := missingNumber([]int{0})
	if actual != 1 {
		t.Fatalf("%d", actual)
	}
}
