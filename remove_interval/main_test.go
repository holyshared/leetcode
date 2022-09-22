package main

import "testing"

func TestA(t *testing.T) {
	actual := removeInterval([][]int{{0, 2}, {3, 4}, {5, 7}}, []int{1, 6})
	t.Fatalf("actual = %v", actual)
}

func TestB(t *testing.T) {
	actual := removeInterval([][]int{{0, 5}}, []int{2, 3})
	t.Fatalf("actual = %v", actual)
}

func TestC(t *testing.T) {
	actual := removeInterval([][]int{{-5, -4}, {-3, -2}, {1, 2}, {3, 5}, {8, 9}}, []int{-1, 4})
	t.Fatalf("actual = %v", actual)
}

func TestD(t *testing.T) {
	actual := removeInterval([][]int{{0, 100}}, []int{0, 50})
	t.Fatalf("actual = %v", actual)
}

func TestE(t *testing.T) {
	actual := removeInterval([][]int{{0, 100}}, []int{50, 100})
	t.Fatalf("actual = %v", actual)
}
