package main

import "testing"

func expect(actual, expected [][]int) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i, v := range actual {
		if v[0] != expected[i][0] || v[1] != expected[i][1] {
			return false
		}
	}
	return true
}

func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 13,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 14,
				},
			},
		},
	}

	actual := closestNodes(root, []int{2, 5, 16})
	expected := [][]int{{2, 2}, {4, 6}, {15, -1}}
	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}

func TestB(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Right: &TreeNode{
			Val: 9,
		},
	}

	actual := closestNodes(root, []int{3})
	expected := [][]int{{-1, 4}}
	if !expect(actual, expected) {
		t.Fatalf("actual = %v, expected = %v", actual, expected)
	}
}
