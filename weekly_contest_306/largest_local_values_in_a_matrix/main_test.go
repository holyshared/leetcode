package main

import (
	"testing"
	"fmt"
)

func TestA(t *testing.T) {
	actual := largestLocal([][]int{
		{9,9,8,1},{5,6,2,6},{8,2,6,4},{6,2,2,2},
	})

	if actual[0][0] != 9 {
		t.Fatalf("actual = %d, expected = 9", actual[0][0])
	}
	if actual[0][1] != 9 {
		t.Fatalf("actual = %d, expected = 9", actual[0][0])
	}
	if actual[1][0] != 8 {
		t.Fatalf("actual = %d, expected = 8", actual[0][0])
	}
	if actual[1][1] != 6 {
		t.Fatalf("actual = %d, expected = v", actual[0][0])
	}
}

func TestB(t *testing.T) {
	actual := largestLocal([][]int{
		{1,1,1,1,1},{1,1,1,1,1},{1,1,2,1,1},{1,1,1,1,1},{1,1,1,1,1},})


		fmt.Println(actual)

	if actual[0][0] != 2 {
		t.Fatalf("actual = %d, expected = 9", actual[0][0])
	}
	if actual[0][1] != 2 {
		t.Fatalf("actual = %d, expected = 9", actual[0][0])
	}
	if actual[1][0] != 2 {
		t.Fatalf("actual = %d, expected = 8", actual[0][0])
	}
	if actual[1][1] != 2 {
		t.Fatalf("actual = %d, expected = v", actual[0][0])
	}
}
