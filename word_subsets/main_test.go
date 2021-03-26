package main

import (
	"testing"
)

func eq(a []string, b []string) bool {
	r := map[string]bool{}
	for _, kw := range a {
		r[kw] = true
	}

	for _, kw := range b {
		_, ok := r[kw]
		if !ok {
			return false
		}
	}
	return true
}

func TestA(t *testing.T) {
	actual := wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"})
	expected := []string{"facebook", "google", "leetcode"}

	if len(actual) != len(expected) {
		t.Fatal("oops!!")
	}

	if !eq(actual, expected) {
		t.Fatal("oops!!")
	}
}

func TestB(t *testing.T) {
	actual := wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"})
	expected := []string{"apple", "google", "leetcode"}

	if len(actual) != len(expected) {
		t.Fatal("oops!!")
	}

	if !eq(actual, expected) {
		t.Fatal("oops!!")
	}
}

func TestC(t *testing.T) {
	actual := wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"})
	expected := []string{"facebook", "google"}

	if len(actual) != len(expected) {
		t.Fatal("oops!!")
	}

	if !eq(actual, expected) {
		t.Fatal("oops!!")
	}
}

func TestD(t *testing.T) {
	actual := wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lo", "eo"})
	expected := []string{"google", "leetcode"}

	if len(actual) != len(expected) {
		t.Fatal("oops!!")
	}

	if !eq(actual, expected) {
		t.Fatal("oops!!")
	}
}

func TestE(t *testing.T) {
	actual := wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"ec", "oc", "ceo"})
	expected := []string{"facebook", "leetcode"}

	if len(actual) != len(expected) {
		t.Fatal("oops!!")
	}

	if !eq(actual, expected) {
		t.Fatal("oops!!")
	}
}
