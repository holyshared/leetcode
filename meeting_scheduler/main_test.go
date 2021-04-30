package main

import "testing"

func TestA(t *testing.T) {
	actual := minAvailableDuration(
		[][]int{
			{10, 50}, {60, 120}, {140, 210},
		},
		[][]int{
			{0, 15}, {60, 70},
		},
		8,
	)
	if actual[0] != 60 {
		t.Fatalf("%d", actual[0])
	}
	if actual[1] != 68 {
		t.Fatalf("%d", actual[1])
	}
}

func TestB(t *testing.T) {
	actual := minAvailableDuration(
		[][]int{
			{10, 50}, {60, 120}, {140, 210},
		},
		[][]int{
			{0, 15}, {60, 70},
		},
		12,
	)
	if len(actual) != 0 {
		t.Fatal("oops!!")
	}
}
