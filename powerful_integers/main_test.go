package main

import (
	"testing"
)

func TestA(t *testing.T) {
	actual := powerfulIntegers(2, 3, 10)
	expected := []int{2, 3, 4, 5, 7, 9, 10}

	for i, v := range actual {
		if expected[i] != v {
			t.Fatalf("%d != %d", expected[i], v)
		}
	}
}

func TestB(t *testing.T) {
	actual := powerfulIntegers(3, 5, 15)
	expected := []int{2, 4, 6, 8, 10, 14}

	for i, v := range actual {
		if expected[i] != v {
			t.Fatalf("%d != %d", expected[i], v)
		}
	}
}
