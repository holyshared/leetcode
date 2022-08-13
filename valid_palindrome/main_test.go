package main

import "testing"

func TestA(t *testing.T) {
	if !isPalindrome("A man, a plan, a canal: Panama") {
		t.Fatalf("acutal = false, expected = true")
	}
}
func TestB(t *testing.T) {
	if !isPalindrome(" ") {
		t.Fatalf("acutal = false, expected = true")
	}
}

func TestC(t *testing.T) {
	if isPalindrome("race a car") {
		t.Fatalf("acutal = true, expected = false")
	}
}

func TestD(t *testing.T) {
	if isPalindrome("0P") {
		t.Fatalf("acutal = true, expected = false")
	}
}

func TestE(t *testing.T) {
	if !isPalindrome("Marge, let's \"[went].\" I await {news} telegram.") {
		t.Fatalf("acutal = false, expected = true")
	}
}
