package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	actual := pathSum(root, 22)

	t.Fatalf("actual = %v", actual)
}

func TestB(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	actual := pathSum(root, 5)

	t.Fatalf("actual = %v", actual)
}

func TestC(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	actual := pathSum(root, 0)

	t.Fatalf("actual = %v", actual)
}

func TestD(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	actual := pathSum(root, 1)

	t.Fatalf("actual = %v", actual)
}

func TestE(t *testing.T) {
	root := &TreeNode{
		Val: -2,
		Right: &TreeNode{
			Val: -3,
		},
	}
	actual := pathSum(root, -5)

	t.Fatalf("actual = %v", actual)
}
