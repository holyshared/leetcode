package main

import "testing"



func Test2(t *testing.T) {
	if isStrobogrammatic("2") != false {
		t.Fatal("oops!!")
	}
}

func Test88(t *testing.T) {
	if !isStrobogrammatic("88") {
		t.Fatal("oops!!")
	}
}


func Test11(t *testing.T) {
	if !isStrobogrammatic("11") {
		t.Fatal("oops!!")
	}
}

func Test962(t *testing.T) {
	if isStrobogrammatic("962") != false {
		t.Fatal("oops!!")
	}
}

func Test69(t *testing.T) {
	if !isStrobogrammatic("69") {
		t.Fatal("oops!!")
	}
}
