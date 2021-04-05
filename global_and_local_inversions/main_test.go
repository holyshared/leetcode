package main

import "testing"

func TestA(t *testing.T) {
	if isIdealPermutation([]int{1, 0, 2}) != true {
		t.Fatal("oops")
	}
}

func TestB(t *testing.T) {
	if isIdealPermutation([]int{1, 2, 0}) != false {
		t.Fatal("oops")
	}
}
