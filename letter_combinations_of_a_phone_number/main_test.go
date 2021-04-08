package main

import "testing"

func TestA(t *testing.T) {
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	actual := letterCombinations("23")

	for i, av := range actual {
		if av != expected[i] {
			t.Fatalf("%s != %s", av, expected[i])
		}
	}
}
