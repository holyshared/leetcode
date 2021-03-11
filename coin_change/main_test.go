package main

import "testing"

func Test125_11(t *testing.T) {
	actual := coinChange([]int{1, 2, 5}, 11)
	if actual != 3 {
		t.Fatalf("%d", actual)
	}
}

func Test2_3(t *testing.T) {
	actual := coinChange([]int{2}, 3)
	if actual != -1 {
		t.Fatalf("%d", actual)
	}
}

func Test1_0(t *testing.T) {
	actual := coinChange([]int{1}, 0)
	if actual != 0 {
		t.Fatalf("%d", actual)
	}
}

func Test112_149(t *testing.T) {
	actual := coinChange([]int{112, 149, 215, 496, 482, 436, 144, 397, 500, 189}, 8480)
	if actual != 17 {
		t.Fatalf("%d", actual)
	}
}
