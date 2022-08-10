package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				} ,
			} ,
		},
	}

	actual := minDepth(root)
	if actual != 5 {
		t.Fatalf("actual = %d", actual)
	}
}