package main

import "testing"

func TestStepAbcdefghijklmnopqrstuvwxyzCba(t *testing.T) {
	actual := calculateTime("abcdefghijklmnopqrstuvwxyz", "cba")

	if actual != 4 {
		t.Fatalf("fatal: %v", actual)
	}
}

func TestStepPqrstuvwxyzabcdefghijklmnoLeetcode(t *testing.T) {
	actual := calculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode")

	if actual != 73 {
		t.Fatalf("fatal: %v", actual)
	}
}
