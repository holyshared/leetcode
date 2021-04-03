package main

import "testing"

func TestA(t *testing.T) {
	actual := longestValidParentheses("(()")
	if actual != 2 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestB(t *testing.T) {
	actual := longestValidParentheses(")()())")
	if actual != 4 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestC(t *testing.T) {
	actual := longestValidParentheses("")
	if actual != 0 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestD(t *testing.T) {
	actual := longestValidParentheses("()())")
	if actual != 4 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestE(t *testing.T) {
	actual := longestValidParentheses("()())()()()()")
	if actual != 8 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestF(t *testing.T) {
	actual := longestValidParentheses("()(())")
	if actual != 6 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestG(t *testing.T) {
	actual := longestValidParentheses("()(()")
	if actual != 2 {
		t.Fatalf("actual: %d", actual)
	}
}
