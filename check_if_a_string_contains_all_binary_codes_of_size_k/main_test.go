package main

import "testing"

func TestA(t *testing.T) {
	actual := hasAllCodes("00110110", 3)
	if actual != true {
		t.Fatal("oops!!")
	}
}
