package main

import (
	"testing"
)

func Test2(t *testing.T) {
	actual := climbStairs(2)
	if actual != 2 {
		t.Fatalf("%d != 2", actual)
	}
}

func Test3(t *testing.T) {
	actual := climbStairs(3)
	if actual != 3 {
		t.Fatalf("%d != 3", actual)
	}
}

func Test5(t *testing.T) {
	actual := climbStairs(5)
	if actual != 8 {
		t.Fatalf("%d != 8", actual)
	}
}

func Test7(t *testing.T) {
	actual := climbStairs(7)
	if actual != 21 {
		t.Fatalf("%d != 21", actual)
	}
}

func Test45(t *testing.T) {
	actual := climbStairs(45)
	if actual != 1836311903 {
		t.Fatalf("%d != 1836311903", actual)
	}
}
