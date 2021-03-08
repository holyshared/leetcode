package main

import "testing"

func TestAba(t *testing.T) {
	if removePalindromeSub("aba") != 1 {
		t.Fatal("oops!!")
	}
}

func TestAbabbabaaba(t *testing.T) {
	if removePalindromeSub("ababbabaaba") != 2 {
		t.Fatal("oops!!")
	}
}
