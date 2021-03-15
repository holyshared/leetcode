package main

import "testing"

func TestEncDec(t *testing.T) {
	codec := Constructor()

	s := codec.Encode("https://leetcode.com/problems/design-tinyurl")
	original := codec.Decode(s)

	if original != "https://leetcode.com/problems/design-tinyurl" {
		t.Fatal("oops!!")
	}
}
