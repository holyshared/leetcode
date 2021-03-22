package main

import (
	"testing"
)

func TestExample1(t *testing.T) {
	wordlist := []string{"KiTe", "kite", "hare", "Hare"}
	queries := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}

	expected := []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"}

	actual := spellchecker(wordlist, queries)

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%d: %s != %s", i, actual[i], expected[i])
		}
	}
}
