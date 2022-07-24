package main

import "testing"

var matrix = [][]int{
	{1, 4, 7, 11, 15},
	{2, 5, 8, 12, 19},
	{3, 6, 9, 16, 22},
	{10, 13, 14, 17, 24},
	{18, 21, 23, 26, 30},
}

func TestA(t *testing.T) {
	actual := searchMatrix(matrix, 5)

	if !actual {
		t.Fatalf("actual = %v", actual)
	}
}

func TestB(t *testing.T) {
	actual := searchMatrix(matrix, 31)

	if actual {
		t.Fatalf("actual = %v", actual)
	}
}

func TestC(t *testing.T) {
	actual := searchMatrix(matrix, 0)

	if actual {
		t.Fatalf("actual = %v", actual)
	}
}

func TestD(t *testing.T) {
	actual := searchMatrix(matrix, 18)

	if !actual {
		t.Fatalf("actual = %v", actual)
	}
}
