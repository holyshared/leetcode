package main

import "testing"

func TestA(t *testing.T) {

	results := runningSum([]int{1, 2, 3, 4})
	expected := []int{1, 3, 6, 10}

	for i, v := range results {
		if expected[i] != v {
			t.Fatalf("index %d: = %d, %d", i, expected, v)
		}
	}

}

func TestB(t *testing.T) {

	results := runningSum([]int{1, 1, 1, 1, 1})
	expected := []int{1, 2, 3, 4, 5}

	for i, v := range results {
		if expected[i] != v {
			t.Fatalf("index %d: = %d, %d", i, expected, v)
		}
	}

}

func TestC(t *testing.T) {

	results := runningSum([]int{3, 1, 2, 10, 1})
	expected := []int{3, 4, 6, 16, 17}

	for i, v := range results {
		if expected[i] != v {
			t.Fatalf("index %d: = %d, %d", i, expected, v)
		}
	}

}
