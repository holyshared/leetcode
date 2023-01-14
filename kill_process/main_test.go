package main

import "testing"

func TestA(t *testing.T) {
	actual := killProcess([]int{1,3,10,5}, []int{3,0,5,3}, 5)
	t.Fatalf("actual = %v", actual)
}

func TestB(t *testing.T) {
	actual := killProcess([]int{1}, []int{0}, 1)
	t.Fatalf("actual = %v", actual)
}

func TestC(t *testing.T) {
	actual := killProcess([]int{1,3,10,5}, []int{3,0,5,3}, 3)
	t.Fatalf("actual = %v", actual)
}