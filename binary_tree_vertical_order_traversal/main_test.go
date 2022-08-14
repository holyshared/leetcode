package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:9,
		},
		Right: &TreeNode{
			Val:20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val:17,
			},
		},
	}
	actual := verticalOrder(root)

	t.Fatalf("%v", actual)
}
