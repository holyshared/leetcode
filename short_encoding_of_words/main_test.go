package main

import "testing"

func TestTimeAndMeAndBell(t *testing.T) {
	refLen := minimumLengthEncoding([]string{"time", "me", "bell"})
	if refLen != 10 {
		t.Fatalf("%d", refLen)
	}
}

func TestT(t *testing.T) {
	refLen := minimumLengthEncoding([]string{"t"})
	if refLen != 2 {
		t.Fatalf("%d", refLen)
	}
}

func TestTimeTimeTimeTimeTime(t *testing.T) {
	refLen := minimumLengthEncoding([]string{"time", "time", "time", "time", "time"})
	if refLen != 5 {
		t.Fatalf("%d", refLen)
	}
}
