package main

import (
	"testing"
)

var root = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 6},
		Right: &TreeNode{Val: 2,

			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 4},
		},
	},
	Right: &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 8},
	},
}

func Test35(t *testing.T) {
	actual := lowestCommonAncestor(root, root, root.Left)

	if actual != root {
		t.Fatalf("%d", actual.Val)
	}
}

func Test51(t *testing.T) {
	actual := lowestCommonAncestor(root, root.Left, root.Right)

	if actual != root {
		t.Fatalf("%d", actual.Val)
	}
}

func Test67(t *testing.T) {
	actual := lowestCommonAncestor(root, root.Left.Left, root.Left.Right.Left)

	if actual != root.Left {
		t.Fatalf("%d", actual.Val)
	}
}

func Test54(t *testing.T) {
	actual := lowestCommonAncestor(root, root.Left, root.Left.Right.Right)

	if actual != root.Left {
		t.Fatalf("%d", actual.Val)
	}
}
