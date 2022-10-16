package main

import "testing"

func TestA(t *testing.T) {
	// 05:00、15:00
	actual := countTime("?5:00")
	if actual != 2 {
		t.Fatalf("actual = %v, expected = 2", actual)
	}
}

func TestB(t *testing.T) {
	actual := countTime("0?:0?")
	if actual != 100 {
		t.Fatalf("actual = %v, expected = 100", actual)
	}
}

func TestC(t *testing.T) {
	actual := countTime("??:??")
	if actual != 1440 {
		t.Fatalf("actual = %v, expected = 1440", actual)
	}
}

func TestD(t *testing.T) {
	// 02:16, 12:16, 22:16
	actual := countTime("?2:16")
	if actual != 3 {
		t.Fatalf("actual = %v, expected = 3", actual)
	}
}

// 20 21 22 23
// 6 * 10
func TestE(t *testing.T) {
	actual := countTime("2?:??")
	if actual != 240 {
		t.Fatalf("actual = %v, expected = 240", actual)
	}
}

func TestF(t *testing.T) {
	actual := countTime("?4:22")
	if actual != 2 {
		t.Fatalf("actual = %v, expected = 3", actual)
	}
}
