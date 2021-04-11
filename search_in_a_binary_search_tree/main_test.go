package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}
	actual := searchBST(root, 2)

	if actual.Val != 2 {
		t.Fatalf("oops!!: %d", actual.Val)
	}

}

func TestB(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}
	actual := searchBST(root, 5)

	if actual != nil {
		t.Fatal("oops!!")
	}

}
