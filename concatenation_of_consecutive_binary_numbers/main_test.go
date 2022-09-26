package main

import "testing"

func TestA(t *testing.T) {
	actual := concatenatedBinary(3)
	if actual != 27 {
		t.Fatalf("actual = %v, expected = 27", actual)
	}
}

func TestB(t *testing.T) {
	actual := concatenatedBinary(12)
	if actual != 505379714 {
		t.Fatalf("actual = %v, expected = 505379714", actual)
	}
}
