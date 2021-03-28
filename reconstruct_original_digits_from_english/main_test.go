package main

import "testing"

func TestA(t *testing.T) {
	actual := originalDigits("owoztneoer")
	if actual != "012" {
		t.Fatal("oops!!")
	}
}

func TestB(t *testing.T) {
	actual := originalDigits("fviefuro")
	if actual != "45" {
		t.Fatalf("oops!! %s", actual)
	}
}

func TestC(t *testing.T) {
	actual := originalDigits("zerozero")
	if actual != "00" {
		t.Fatalf("oops!! %s", actual)
	}
}

func TestD(t *testing.T) {
	actual := originalDigits("zeroonetwothreefourfivesixseveneightnine")
	if actual != "0123456789" {
		t.Fatalf("oops!! %s", actual)
	}
}
