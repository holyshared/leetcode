package main

import "testing"

func TestA(t *testing.T) {
	actual := evalRPN([]string{"2", "1", "+", "3", "*"})
	if actual != 9 {
		t.Fatalf("actual = %v, expected = %v", actual, 9)
	}
}

func TestB(t *testing.T) {
	actual := evalRPN([]string{"4", "13", "5", "/", "+"})
	if actual != 6 {
		t.Fatalf("actual = %v, expected = %v", actual, 6)
	}
}

func TestC(t *testing.T) {
	actual := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	if actual != 22 {
		t.Fatalf("actual = %v, expected = %v", actual, 22)
	}
}
