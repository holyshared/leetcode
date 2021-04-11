package main

import "testing"

func TestA(t *testing.T) {
	actual := getRow(3)
	expected := []int{1, 3, 3, 1}

	for i, v := range actual {
		if v != expected[i] {
			t.Fatalf("%d != %d", v, expected[i])
		}
	}
}

func TestB(t *testing.T) {
	actual := getRow(4)
	expected := []int{1, 4, 6, 4, 1}

	for i, v := range actual {
		if v != expected[i] {
			t.Fatalf("%d != %d", v, expected[i])
		}
	}
}

func TestC(t *testing.T) {
	actual := getRow(5)
	expected := []int{1, 5, 10, 10, 5, 1}

	for i, v := range actual {
		if v != expected[i] {
			t.Fatalf("%d != %d", v, expected[i])
		}
	}
}
