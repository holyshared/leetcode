package main

import "testing"

func expect(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestA(t *testing.T) {
	expected := []string{"2", "4->49", "51->74", "76->99"}
	actual := findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99)

	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}

func TestB(t *testing.T) {
	expected := []string{"-10->-1", "2", "4->49", "51->74"}
	actual := findMissingRanges([]int{0, 1, 3, 50, 75}, -10, 75)

	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}

func TestC(t *testing.T) {
	expected := []string{"1"}
	actual := findMissingRanges([]int{}, 1, 1)

	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}

func TestD(t *testing.T) {
	expected := []string{"-1"}
	actual := findMissingRanges([]int{}, -1, -1)

	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}

func TestE(t *testing.T) {
	expected := []string{"-3->-1"}
	actual := findMissingRanges([]int{}, -3, -1)

	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}

func TestF(t *testing.T) {
	expected := []string{"-2"}
	actual := findMissingRanges([]int{-1}, -2, -1)

	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}

func TestG(t *testing.T) {
	expected := []string{"-2", "0"}
	actual := findMissingRanges([]int{-1}, -2, 0)

	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}
