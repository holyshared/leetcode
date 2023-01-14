package main

import "testing"

func TestA(t *testing.T) {
	actual := pushDominoes("RR.L")
	t.Fatalf("actual = %v", actual)
}

func TestR(t *testing.T) {
	actual := pushDominoes("R..L")
	t.Fatalf("actual = %v", actual)
}