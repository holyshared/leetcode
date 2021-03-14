package main

import "testing"

func Test122(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}

	if isSymmetric(root) != true {
		t.Fatal("oops!!")
	}
}

/*
*                     1
*              2            2
*           3     4       4      3
*         5   6 7   8   8   7  6    5

2 3 5 6 4 7 8

2 3 6 5 4 8 7
*/

func Test122__(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 6},
			},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 8},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 5},
			},
		},
	}

	if isSymmetric(root) != true {
		t.Fatal("oops!!")
	}
}
