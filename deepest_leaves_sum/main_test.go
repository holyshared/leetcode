package main

import "testing"

func TestA(t *testing.T) {
	actual := deepestLeavesSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 7}},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 8}},
		},
	})

	if actual != 15 {
		t.Fatalf("%d", actual)
	}
}
