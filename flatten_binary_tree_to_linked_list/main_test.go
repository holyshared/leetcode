package main

import (
	"testing"
)

func nodeToList(node *TreeNode) []int{
	result := []int{}

	for node != nil {
		result = append(result, node.Val)
		node = node.Right
	}
	return result
}

func allLeftNodeIsNil(node *TreeNode) (bool, int) {
	for node != nil {
		if node.Left != nil {
			return false, node.Val
		}
		node = node.Right
	}
	return true, -1
}

func expectList(a, b []int) bool {
	for i, val := range a {
		if b[i] != val {
			return false
		}
	}
	return true
}


func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	flatten(root)

	actual := nodeToList(root)
	expected := []int{1,2,3,4,5,6}

	if !expectList(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}

	if ok, val := allLeftNodeIsNil(root); !ok {
		t.Fatalf("node: %v left is not nil", val)
	}
}

func TestB(t *testing.T) {
	root := &TreeNode{ Val: 0 }
	flatten(root)

	if root == nil {
		t.Fatal("actual = nil, expected = 1")
	}
}