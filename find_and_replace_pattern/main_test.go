package main

import "testing"

func TestA(t *testing.T) {
	actual := findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb")

	if len(actual) != 2 {
		t.Fatalf("actual = %v", actual)
	}

}
