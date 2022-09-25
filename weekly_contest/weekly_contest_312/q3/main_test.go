package main

import "testing"

func TestA(t *testing.T) {
	actual := goodIndices([]int{2, 1, 1, 1, 3, 4, 1}, 2)

	if actual[0] != 2 {
		t.Fatalf("actual = %v", actual)
	}
	if actual[1] != 3 {
		t.Fatalf("actual = %v", actual)
	}
}

func TestB(t *testing.T) {
	actual := goodIndices([]int{2, 1, 1, 2}, 2)

	if len(actual) != 0 {
		t.Fatalf("actual = %v", actual)
	}
}

func TestC(t *testing.T) {
	actual := goodIndices([]int{440043, 276285, 336957}, 1)

	if len(actual) != 1 {
		t.Fatalf("actual = %v", actual)
	}
}

func TestD(t *testing.T) {
	actual := goodIndices([]int{388589, 17165, 726687, 401298, 600033, 537254, 301052, 151069, 399955}, 4)

	if len(actual) != 0 {
		t.Fatalf("actual = %v", actual)
	}
}
