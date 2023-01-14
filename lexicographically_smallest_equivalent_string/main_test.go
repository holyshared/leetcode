package main

import "testing"

// s1 = "parker", s2 = "morris", baseStr = "parser"
func TestA(t *testing.T) {
	actual := smallestEquivalentString("parker", "morris", "parser")
	t.Fatalf("actual = %v", actual)
}
