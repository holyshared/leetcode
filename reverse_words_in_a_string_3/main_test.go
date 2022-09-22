package main

import "testing"

func TestA(t *testing.T) {
	actual := reverseWords("Let's take LeetCode contest")
	expected := "s'teL ekat edoCteeL tsetnoc"
	if actual != expected {
		t.Fatalf("actual = %s, expected = %s", actual, expected)
	}
}

func TestB(t *testing.T) {
	actual := reverseWords("God Ding")
	expected := "doG gniD"
	if actual != expected {
		t.Fatalf("actual = %s, expected = %s", actual, expected)
	}
}

func TestC(t *testing.T) {
	actual := reverseWords("God")
	expected := "doG"
	if actual != expected {
		t.Fatalf("actual = %s, expected = %s", actual, expected)
	}
}
