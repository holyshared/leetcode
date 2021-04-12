package main

import "testing"

func TestA(t *testing.T) {
	actual := kthGrammar(4, 5)
	if actual != 1 {
		t.Fatalf("oops!: %d", actual)
	}
}
