package main

import "testing"

func TestN3(t *testing.T) {
	actual := minOperations(3)
	if actual != 2 {
		t.Fatalf("%d", actual)
	}
}

func TestN5(t *testing.T) {
	actual := minOperations(5)
	if actual != 6 {
		t.Fatalf("%d", actual)
	}
}

func TestN6(t *testing.T) {
	actual := minOperations(6)
	if actual != 9 {
		t.Fatalf("%d", actual)
	}
}
