package main

import "testing"

func TestWdcab_1(t *testing.T) {
	actual := smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}})
	if actual != "bacd" {
		t.Fatalf("actual = %s, expected = bacd", actual)
	}
}

func TestWdcab_2(t *testing.T) {
	actual := smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0,2}})
	if actual != "abcd" {
		t.Fatalf("actual = %s, expected = abcd", actual)
	}
}
