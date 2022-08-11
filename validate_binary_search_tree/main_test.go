package main

import "testing"
/*
func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	actual := isValidBST(root)
	if !actual {
		t.Fatalf("actual = %v, expected = true", actual)
	}
}

func TestB(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	actual := isValidBST(root)
	if actual {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}
*/
func TestC(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	actual := isValidBST(root)
	if actual {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}
