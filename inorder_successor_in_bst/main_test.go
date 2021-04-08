package main

import "testing"

func TestA(t *testing.T) {
	p := &TreeNode{Val: 1}
	root := &TreeNode{Val: 2, Left: p, Right: &TreeNode{Val: 3}}

	actual := inorderSuccessor(root, p)
	if actual.Val != 2 {
		t.Fatal("oops!!")
	}
}

func TestB(t *testing.T) {
	p := &TreeNode{Val: 6}
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: p,
	}

	actual := inorderSuccessor(root, p)
	if actual != nil {
		t.Fatal("oops!!")
	}
}

func TestC(t *testing.T) {
	p := &TreeNode{Val: 3}
	root := &TreeNode{
		Val:   2,
		Right: p,
	}

	actual := inorderSuccessor(root, root)
	if actual.Val != 3 {
		t.Fatal("oops!!")
	}
}
