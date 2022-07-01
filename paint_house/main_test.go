package main

import "testing"

func TestP1(t *testing.T) {
	consts := [][]int{{20,18,4},{9,9,10}}

	actual := minCost(consts)
	if actual != 13 {
		t.Fatalf("oops!! actual = %d", actual)
	}
}


func TestP2(t *testing.T) {
	consts := [][]int{
		{5,8,6},{19,14,13},{7,5,12},{14,15,17},{3,20,10},
	}

	actual := minCost(consts)
	if actual != 43 {
		t.Fatalf("oops!! actual = %d", actual)
	}
}
