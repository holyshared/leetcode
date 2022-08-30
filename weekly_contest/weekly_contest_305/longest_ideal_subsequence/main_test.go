package main

import "testing"

func TestA(t *testing.T) {
	actual := longestIdealString("acfgbd", 2)
	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}

func TestB(t *testing.T) {
	actual := longestIdealString("abcd", 3)
	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}

func TestC(t *testing.T) {
	actual := longestIdealString("pvjcci", 4)
	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}
