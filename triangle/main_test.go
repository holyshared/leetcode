package main

import "testing"

func TestSum11(t *testing.T) {
	actual := minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	})
	if actual != 11 {
		t.Fatalf("%d", actual)
	}
}

func TestSumMinus10(t *testing.T) {
	actual := minimumTotal([][]int{
		{-10},
	})
	if actual != -10 {
		t.Fatalf("%d", actual)
	}
}

func TestC(t *testing.T) {
	actual := minimumTotal([][]int{
		{-1},
		{2, 3},
		{1, -1, -3},
	})
	if actual != -1 {
		t.Fatalf("%d", actual)
	}
}
