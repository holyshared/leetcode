package main

import "testing"

func Test112233(t *testing.T) {
	actual := distributeCandies([]int{1, 1, 2, 2, 3, 3})
	if actual != 3 {
		t.Fatalf("%d", actual)
	}
}

func Test1123(t *testing.T) {
	actual := distributeCandies([]int{1, 1, 2, 3})
	if actual != 2 {
		t.Fatalf("%d", actual)
	}
}

func Test6666(t *testing.T) {
	actual := distributeCandies([]int{6, 6, 6, 6})
	if actual != 1 {
		t.Fatalf("%d", actual)
	}
}
