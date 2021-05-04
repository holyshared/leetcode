package main

import "testing"

func TestA(t *testing.T) {
	acutal := checkPossibility([]int{4, 2, 3})

	if acutal != true {
		t.Fatal("oops!!")
	}
}

func TestB(t *testing.T) {
	acutal := checkPossibility([]int{4, 2, 1})

	if acutal != false {
		t.Fatal("oops!!")
	}
}

func TestC(t *testing.T) {
	acutal := checkPossibility([]int{2, 2, 3})

	if acutal != true {
		t.Fatal("oops!!")
	}
}

func TestD(t *testing.T) {
	acutal := checkPossibility([]int{1, 2, 3})

	if acutal != true {
		t.Fatal("oops!!")
	}
}

func TestE(t *testing.T) {
	acutal := checkPossibility([]int{3, 4, 2, 3})

	if acutal != false {
		t.Fatal("oops!!")
	}
}
