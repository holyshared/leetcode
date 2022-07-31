package main

import "testing"

func TestA(t *testing.T) {
	numArray := Constructor([]int{1, 3, 5})

	if sum := numArray.SumRange(0, 2); sum != 9 {
		t.Fatalf("actual = %d, expected = 9", sum)
	}

	numArray.Update(1, 2)

	if sum := numArray.SumRange(0, 2); sum != 8 {
		t.Fatalf("actual = %d, expected = 8", sum)
	}
}

func TestB(t *testing.T) {
	numArray := Constructor([]int{1, 3})

	if sum := numArray.SumRange(0, 1); sum != 4 {
		t.Fatalf("actual = %d, expected = 4", sum)
	}

	numArray.Update(1, 2)

	if sum := numArray.SumRange(0, 1); sum != 3 {
		t.Fatalf("actual = %d, expected = 3", sum)
	}
}

func TestC(t *testing.T) {
	numArray := Constructor([]int{1})

	if sum := numArray.SumRange(0, 0); sum != 1 {
		t.Fatalf("actual = %d, expected = 1", sum)
	}
}