package main

import "testing"

func TestA(t *testing.T) {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}
	actual := connect(root)

	if actual.Val != 1 {
		t.Fatalf("actual = %d, expected = 1", actual.Val)
	}
}
