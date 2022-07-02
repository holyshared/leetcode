package main

import (
	"testing"
)

func TestA(t *testing.T) {
	equations := [][]string{

		{"a", "b"}, {"b", "c"},
	}
	values := []float64{2.0, 3.0}
	queries := [][]string{
		{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"},
	}

	actual := calcEquation(equations, values, queries)

	if actual[0] != 6.0 {
		t.Fatalf("actual = %f., expected = 6.0", actual[0])
	}
	if actual[1] != 0.5 {
		t.Fatalf("actual = %f., expected = 0.5", actual[1])
	}
}
