package main

import "testing"

func TestA(t *testing.T) {
	actual := minCostClimbingStairs([]int{10,15,20})
	if actual != 15 {
		t.Fatalf("actual = %d, expected = 15", actual)
	}
}

func TestB(t *testing.T) {
	actual := minCostClimbingStairs([]int{1,100,1,1,1,100,1,1,100,1})
	if actual != 6 {
		t.Fatalf("actual = %d, expected = 6", actual)
	}
}