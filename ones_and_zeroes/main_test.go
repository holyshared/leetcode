package main

import "testing"

func TestA(t *testing.T) {
	actual := findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)
	if actual != 4 {
		t.Fatalf("%d", actual)
	}

}

func TestB(t *testing.T) {
	actual := findMaxForm([]string{"10", "1", "0"}, 1, 1)
	if actual != 2 {
		t.Fatalf("%d", actual)
	}
}
