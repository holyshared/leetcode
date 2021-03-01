package main

import "testing"

func Test321(t *testing.T) {
	actual := thirdMax([]int{3, 2, 1})
	if actual != 1 {
		t.Fatalf("%d", actual)
	}
}

func Test12(t *testing.T) {
	actual := thirdMax([]int{1, 2})
	if actual != 2 {
		t.Fatalf("%d", actual)
	}
}

func Test2231(t *testing.T) {
	actual := thirdMax([]int{2, 2, 3, 1})
	if actual != 1 {
		t.Fatalf("%d", actual)
	}
}

func Test122535(t *testing.T) {
	actual := thirdMax([]int{1, 2, 2, 5, 3, 5})
	if actual != 2 {
		t.Fatalf("%d", actual)
	}
}
