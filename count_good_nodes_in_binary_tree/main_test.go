package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	actual := goodNodes(root)
	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
}

func TestB(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	actual := goodNodes(root)
	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}

func TestC(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	actual := goodNodes(root)
	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}
