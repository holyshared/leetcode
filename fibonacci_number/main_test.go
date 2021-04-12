package main

import "testing"

func Test2(t *testing.T) {
	actual := fib(2)
	if actual != 1 {
		t.Fatalf("%d", actual)
	}
}

func Test4(t *testing.T) {
	actual := fib(4)
	if actual != 3 {
		t.Fatalf("%d", actual)
	}
}
