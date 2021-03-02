package main

import "testing"

func Test1224(t *testing.T) {
	actual := findErrorNums([]int{1, 2, 2, 4})

	if actual[0] != 2 {
		t.Fatalf("%d", actual[0])
	}

	if actual[1] != 3 {
		t.Fatalf("%d", actual[1])
	}
}

func Test11(t *testing.T) {

	actual := findErrorNums([]int{1, 1})

	if actual[0] != 1 {
		t.Fatalf("%d", actual[0])
	}

	if actual[1] != 2 {
		t.Fatalf("%d", actual[1])
	}
}

func Test323465(t *testing.T) {
	actual := findErrorNums([]int{3, 2, 3, 4, 6, 5})

	if actual[0] != 3 {
		t.Fatalf("%d", actual[0])
	}

	if actual[1] != 1 {
		t.Fatalf("%d", actual[1])
	}
}

func Test1532276489(t *testing.T) {
	actual := findErrorNums([]int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9})

	if actual[0] != 2 {
		t.Fatalf("%d", actual[0])
	}

	if actual[1] != 10 {
		t.Fatalf("%d", actual[1])
	}
}
