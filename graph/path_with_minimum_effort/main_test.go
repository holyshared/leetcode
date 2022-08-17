package main

import "testing"

func TestA(t *testing.T) {
	heights := [][]int{{1,2,2},{3,8,2},{5,3,5}}
	actual := minimumEffortPath(heights)

	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}

func TestB(t *testing.T) {
	heights := [][]int{
		{1,2,3},
			{3,8,4},
				{5,3,5},
	}
	actual := minimumEffortPath(heights)

	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}

func TestC(t *testing.T) {
	heights := [][]int{
		{1,2,1,1,1},{1,2,1,2,1},{1,2,1,2,1},{1,2,1,2,1},{1,1,1,2,1},
	}
	actual := minimumEffortPath(heights)

	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}

func TestD(t *testing.T) {
	heights := [][]int{
		{3},
	}
	actual := minimumEffortPath(heights)

	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}