package main

import (
	"testing"
)

func TestA(t *testing.T) {
	actual := movesToStamp("abc", "ababc")
	expect := []int{1, 0, 2}

	for i, v := range actual {
		if v != expect[i] {
			t.Fatal("oops!!")
		}
	}
}

func TestB(t *testing.T) {
	actual := movesToStamp("abca", "aabcaca")
	expect := []int{2, 3, 0, 1}

	for i, v := range actual {
		if v != expect[i] {
			t.Fatal("oops!!")
		}
	}
}
