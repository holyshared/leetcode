package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}
	actual := longestConsecutive(root)

	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}

func TestB(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	actual := longestConsecutive(root)

	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}

func TestC(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	actual := longestConsecutive(root)

	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}
