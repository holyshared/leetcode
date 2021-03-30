package main

import "testing"

func TestA(t *testing.T) {
	envelopes := [][]int{
		[]int{5, 4},
		[]int{6, 4},
		[]int{6, 7},
		[]int{2, 3},
	}
	actual := maxEnvelopes(envelopes)
	if actual != 3 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestB(t *testing.T) {
	envelopes := [][]int{
		[]int{1, 1},
		[]int{1, 1},
		[]int{1, 1},
	}
	actual := maxEnvelopes(envelopes)
	if actual != 1 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestC(t *testing.T) {
	envelopes := [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
	}
	actual := maxEnvelopes(envelopes)
	if actual != 3 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestD(t *testing.T) {
	envelopes := [][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{3, 4},
	}
	actual := maxEnvelopes(envelopes)
	if actual != 2 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestE(t *testing.T) {
	envelopes := [][]int{

		[]int{2, 100},
		[]int{3, 200},
		[]int{4, 300},
		[]int{5, 500},
		[]int{5, 400},
		[]int{5, 250},
		[]int{6, 370},
		[]int{6, 360},
		[]int{7, 380},
	}
	actual := maxEnvelopes(envelopes)
	if actual != 5 {
		t.Fatalf("actual: %d", actual)
	}
}

func TestF(t *testing.T) {
	envelopes := [][]int{
		[]int{46, 89},
		[]int{50, 53},
		[]int{52, 68},
		[]int{72, 45},
		[]int{77, 81},
	}
	actual := maxEnvelopes(envelopes)
	if actual != 3 {
		t.Fatalf("actual: %d", actual)
	}
}
