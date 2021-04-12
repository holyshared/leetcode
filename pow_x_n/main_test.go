package main

import (
	"math"
	"testing"
)

func Test2_10(t *testing.T) {
	actual := myPow(2.0, 10)
	if actual != math.Round(1024.0) {
		t.Fatalf("%f", actual)
	}
}

func Test21_2(t *testing.T) {
	actual := myPow(2.1, 3)
	if math.Round(actual) != math.Round(9.261000) {
		t.Fatalf("%f != %f", actual, float64(9.261000))
	}
}

func Test2_m2(t *testing.T) {
	actual := myPow(2.0, -2)
	if math.Round(actual) != math.Round(0.25000) {
		t.Fatalf("%f != %f", actual, float64(0.25000))
	}
}

func Test34_m3(t *testing.T) {
	actual := myPow(34.00515, -3)
	if math.Round(actual) != math.Round(3e-05) {
		t.Fatalf("%f != %f", actual, float64(3e-05))
	}
}
