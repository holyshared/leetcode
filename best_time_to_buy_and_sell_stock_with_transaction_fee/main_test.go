package main

import "testing"

func Test132849(t *testing.T) {
	actual := maxProfit([]int{1,3,2,8,4,9}, 2)

	if actual != 8 {
		t.Fatal("oops!!")
	}
}
