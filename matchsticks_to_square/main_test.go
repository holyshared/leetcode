package main

import "testing"


func TestA(t *testing.T) {
	actual := makesquare([]int{1, 1, 2, 2, 2})
	if actual != true {
		t.Fatalf("actual = %v, expected = true", actual)
	}

}

func TestB(t *testing.T) {
	actual := makesquare([]int{3, 3, 3, 3, 4})
	if actual != false {
		t.Fatalf("actual = %v, expected = false", actual)
	}

}

func TestC(t *testing.T) {
	actual := makesquare([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 4, 3, 2, 1})
	if actual != false {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}
