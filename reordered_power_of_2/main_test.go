package main

import "testing"

func Test1(t *testing.T) {
	if reorderedPowerOf2(1) != true {
		t.Fatal("oops!!")
	}
}

func Test10(t *testing.T) {
	if reorderedPowerOf2(10) != false {
		t.Fatal("oops!!")
	}
}

func Test16(t *testing.T) {
	if reorderedPowerOf2(16) != true {
		t.Fatal("oops!!")
	}
}

func Test46(t *testing.T) {
	if reorderedPowerOf2(46) != true {
		t.Fatal("oops!!")
	}
}
