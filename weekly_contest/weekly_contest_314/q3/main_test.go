package main

import "testing"

func TestA(t *testing.T) {
	actual := robotWithString("zza")
	t.Fatalf("actual = %v", actual)
}

func TestB(t *testing.T) {
	actual := robotWithString("bac")
	t.Fatalf("actual = %v", actual)
}

func TestC(t *testing.T) {
	actual := robotWithString("bdda")
	t.Fatalf("actual = %v", actual)
}

func TestD(t *testing.T) {
	// expected: bdevfziy
	actual := robotWithString("bydizfve")
	t.Fatalf("actual = %v", actual)
}
