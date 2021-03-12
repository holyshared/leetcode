package main

import "testing"

func Test00110110(t *testing.T) {
	actual := hasAllCodes("00110110", 2)
	if actual != true {
		t.Fatal("oops!!")
	}
}

func Test00110(t *testing.T) {
	actual := hasAllCodes("00110", 2)
	if actual != true {
		t.Fatal("oops!!")
	}
}

func Test0110_2(t *testing.T) {
	actual := hasAllCodes("0110", 1)
	if actual != true {
		t.Fatal("oops!!")
	}
}

func Test0110_1(t *testing.T) {
	actual := hasAllCodes("0110", 2)
	if actual != false {
		t.Fatal("oops!!")
	}
}

func Test0000000001011100(t *testing.T) {
	actual := hasAllCodes("0000000001011100", 4)
	if actual != false {
		t.Fatal("oops!!")
	}
}

func Test000000000000000000001011100(t *testing.T) {
	actual := hasAllCodes("000000000000000000001011100", 20)
	if actual != false {
		t.Fatal("oops!!")
	}
}
