package main

import "testing"

func Test3(t *testing.T) {
	actual := intToRoman(3)
	if actual != "III" {
		t.Fatalf("%s", actual)
	}
}

func Test4(t *testing.T) {
	actual := intToRoman(4)
	if actual != "IV" {
		t.Fatalf("%s", actual)
	}
}

func Test9(t *testing.T) {
	actual := intToRoman(9)
	if actual != "IX" {
		t.Fatalf("%s", actual)
	}
}

func Test58(t *testing.T) {
	actual := intToRoman(58)
	if actual != "LVIII" {
		t.Fatalf("%s", actual)
	}
}

func Test1994(t *testing.T) {
	actual := intToRoman(1994)
	if actual != "MCMXCIV" {
		t.Fatalf("%s", actual)
	}
}

func Test10(t *testing.T) {
	actual := intToRoman(10)
	if actual != "X" {
		t.Fatalf("%s", actual)
	}
}

func Test100(t *testing.T) {
	actual := intToRoman(100)
	if actual != "C" {
		t.Fatalf("%s", actual)
	}
}

func Test500(t *testing.T) {
	actual := intToRoman(500)
	if actual != "D" {
		t.Fatalf("%s", actual)
	}
}

func Test1000(t *testing.T) {
	actual := intToRoman(1000)
	if actual != "M" {
		t.Fatalf("%s", actual)
	}
}
