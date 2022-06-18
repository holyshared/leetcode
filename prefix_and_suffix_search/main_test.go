package main

import (
	"testing"
)

func TestWordsApplePSae(t *testing.T) {
	wordFilter := Constructor([]string{"apple"})
	actual := wordFilter.F("a", "e")
	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}
