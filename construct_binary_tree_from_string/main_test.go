package main

import "testing"

func Test423165(t *testing.T) {
	result := str2tree("4(2(3)(1))(6(5))")

	actual := []int{
		result.Val,
		result.Left.Val,
		result.Left.Left.Val,
		result.Left.Right.Val,
		result.Right.Val,
		result.Right.Left.Val,
	}
	expect := []int{
		4,
		2,
		3,
		1,
		6,
		5,
	}

	for i, v := range actual {
		if v != expect[i] {
			t.Fatalf("%d, %d", v, expect[i])
		}
	}
}
