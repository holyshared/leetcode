package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}

	actual := rightSideView(root)

	if actual[0] != 1 {
		t.Fatalf("actual = %d, expected = 1", actual[0])
	}
	if actual[1] != 3 {
		t.Fatalf("actual = %d, expected = 3", actual[1])
	}
	if actual[2] != 4 {
		t.Fatalf("actual = %d, expected = 4", actual[2])
	}

}
