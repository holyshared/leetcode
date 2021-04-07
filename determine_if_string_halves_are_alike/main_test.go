package main

import "testing"

func TestA(t *testing.T) {
	if !halvesAreAlike("book") {
		t.Fatal("oops!!")
	}
}

func TestB(t *testing.T) {
	if halvesAreAlike("textbook") {
		t.Fatal("oops!!")
	}
}

func TestC(t *testing.T) {
	if halvesAreAlike("MerryChristmas") {
		t.Fatal("oops!!")
	}
}

func TestD(t *testing.T) {
	if !halvesAreAlike("AbCdEfGh") {
		t.Fatal("oops!!")
	}
}
