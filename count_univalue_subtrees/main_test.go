package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{ Val: 5, },
			Right: &TreeNode{ Val: 5, },
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{ Val: 5, },
		},
	}
	actual := countUnivalSubtrees(root)

	if actual != 4 {
		t.Fatalf("actual = %d", actual)
	}
}
