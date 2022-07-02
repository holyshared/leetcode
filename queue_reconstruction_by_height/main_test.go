package main

import (
	"testing"
)

func TestA(t *testing.T) {
	people := [][]int{ {7,0},{4,4},{7,1},{5,0},{6,1},{5,2}} 

	actual := reconstructQueue(people)

	expected := [][]int{{5,0},{7,0},{5,2},{6,1},{4,4},{7,1} }

	for i := 0; i< len(actual); i++ {
		if actual[i][0] != expected[i][0] {
			t.Fatalf("actual[%d][0] = %d, expected[%d][0] = %d", i, actual[i][0], i, expected[i][0] )
		}
	}
}