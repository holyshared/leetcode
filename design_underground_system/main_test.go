package main

import "testing"

func TestExample1(t *testing.T) {
	us := Constructor()

	us.CheckIn(45, "Leyton", 3)
	us.CheckIn(32, "Paradise", 8)
	us.CheckIn(27, "Leyton", 10)

	us.CheckOut(45, "Waterloo", 15)
	us.CheckOut(27, "Waterloo", 20)
	us.CheckOut(32, "Cambridge", 22)

	paradiseToCambridge := us.GetAverageTime("Paradise", "Cambridge")
	if paradiseToCambridge != 14.00 {
		t.Fatalf("Paradise to Cambridge %f", paradiseToCambridge)
	}

	leytonToWaterloo := us.GetAverageTime("Leyton", "Waterloo")
	if leytonToWaterloo != 11.00 {
		t.Fatalf("Leyton to Waterloo %f", leytonToWaterloo)
	}

	us.CheckIn(10, "Leyton", 24)

	leytonToWaterloo = us.GetAverageTime("Leyton", "Waterloo")
	if leytonToWaterloo != 11.00 {
		t.Fatalf("Leyton to Waterloo %f", leytonToWaterloo)
	}

	us.CheckOut(10, "Waterloo", 38)

	leytonToWaterloo = us.GetAverageTime("Leyton", "Waterloo")
	if leytonToWaterloo != 12.00 {
		t.Fatalf("Leyton to Waterloo %f", leytonToWaterloo)
	}
}

func TestExample2(t *testing.T) {
	us := Constructor()

	us.CheckIn(10, "Leyton", 3)
	us.CheckOut(10, "Paradise", 8)

	leytonToParadise := us.GetAverageTime("Leyton", "Paradise")
	if leytonToParadise != 5.00000 {
		t.Fatalf("Leyton to Paradise %f", leytonToParadise)
	}

	us.CheckIn(5, "Leyton", 10)
	us.CheckOut(5, "Paradise", 16)
	leytonToParadise = us.GetAverageTime("Leyton", "Paradise")
	if leytonToParadise != 5.50000 {
		t.Fatalf("Leyton to Paradise %f", leytonToParadise)
	}

	us.CheckIn(2, "Leyton", 21)
	us.CheckOut(2, "Paradise", 30)
	leytonToParadise = us.GetAverageTime("Leyton", "Paradise")
	if leytonToParadise <= 6.666667 {
		t.Fatalf("Leyton to Paradise %f", leytonToParadise)
	}
}
