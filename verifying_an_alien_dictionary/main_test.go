package main

import "testing"

func TestA(t *testing.T) {
	if !isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz") {
		t.Fatal("oops!!")
	}
}

func TestB(t *testing.T) {
	if isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz") {
		t.Fatal("oops!!")
	}
}

func TestC(t *testing.T) {
	if isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz") {
		t.Fatal("oops!!")
	}
}

func TestD(t *testing.T) {
	if !isAlienSorted([]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz") {
		t.Fatal("oops!!")
	}
}

func TestE(t *testing.T) {
	if !isAlienSorted([]string{"ubg", "kwh"}, "qcipyamwvdjtesbghlorufnkzx") {
		t.Fatal("oops!!")
	}
}
