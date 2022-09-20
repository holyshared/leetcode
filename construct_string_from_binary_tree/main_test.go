package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left:  &TreeNode{
			Val: 2,
			Left:  &TreeNode{
				Val: 4,
			},
		},
		Right:  &TreeNode{
			Val: 3,
		},
	}

	actual := tree2str(root)

	if actual != "1(2(4))(3)" {
		t.Fatalf("actual = %s, expected = 1(2(4))(3)", actual)
	}
}

func TestB(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left:  &TreeNode{
			Val: 2,
			Right:  &TreeNode{
				Val: 4,
			},
		},
		Right:  &TreeNode{
			Val: 3,
		},
	}

	actual := tree2str(root)

	if actual != "1(2()(4))(3)" {
		t.Fatalf("actual = %s, expected = 1(2()(4))(3)", actual)
	}
}