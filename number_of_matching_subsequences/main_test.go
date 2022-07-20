package main

import "testing"

func TestOne1(t *testing.T) {
	actual := numMatchingSubseq("dsahjpjauf", []string{"ja"})
	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}

func TestOne2(t *testing.T) {
	actual := numMatchingSubseq("abcde", []string{"acd"})
	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}

func TestOneMany(t *testing.T) {
	actual := numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"})
	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}


func TestLong(t *testing.T) {
	actual := numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"})
	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}
