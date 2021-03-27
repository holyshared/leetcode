package main

import "testing"

func TestA(t *testing.T) {
	actual := countSubstrings("abc")
	if actual != 3 {
		t.Fatalf("%d", actual)
	}
}

func TestB(t *testing.T) {
	actual := countSubstrings("aaa")
	if actual != 6 {
		t.Fatalf("%d", actual)
	}
}
