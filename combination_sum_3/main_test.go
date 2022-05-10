package main

import (
	"testing"
)

func Test_k3n7(t *testing.T) {
	actual := combinationSum3(3, 7)

	if len(actual) != 1 {
		t.Fatalf("actual = %d", actual)
	}
}

func Test_k3n9(t *testing.T) {
	actual := combinationSum3(3, 9)
	if len(actual) != 3 {
		t.Fatalf("actual = %d", actual)
	}
}
