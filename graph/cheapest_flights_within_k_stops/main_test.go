package main

import "testing"

func TestA(t *testing.T) {
	flights := [][]int{{0,1,100},{1,2,100},{0,2,500}}

	actual := findCheapestPrice(3, flights, 0, 2, 1)

	if actual != 200 {
		t.Fatalf("actual = %d, expected = 200", actual)
	}
}

func TestB(t *testing.T) {
	flights := [][]int{
		{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200},
	}
	actual := findCheapestPrice(4, flights, 0, 3, 1)

	if actual != 700 {
		t.Fatalf("actual = %d, expected = 700", actual)
	}
}


func TestC(t *testing.T) {
	flights := [][]int{
		{0,1,100},{1,2,100},{0,2,500},
	}
	actual := findCheapestPrice(3, flights, 0, 2, 0)

	if actual != 500 {
		t.Fatalf("actual = %d, expected = 500", actual)
	}
}
