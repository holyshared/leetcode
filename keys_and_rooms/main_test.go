package main

import "testing"

func TestA(t *testing.T) {
	rooms := [][]int{[]int{1}, []int{2}, []int{3}, []int{}}

	if canVisitAllRooms(rooms) != true {
		t.Fatal("oops!!")
	}
}

func TestB(t *testing.T) {
	rooms := [][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{}}

	if canVisitAllRooms(rooms) != false {
		t.Fatal("oops!!")
	}
}

func TestC(t *testing.T) {
	rooms := [][]int{
		[]int{3},
		[]int{2},
		[]int{2},
		[]int{1},
	}

	if canVisitAllRooms(rooms) != true {
		t.Fatal("oops!!")
	}
}
