package main

import "testing"

func Test392015(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	results := averageOfLevels(root)

	expected := []float64{3, 14.5, 11}

	for i := 0; i < len(expected); i++ {
		if expected[i] != results[i] {
			t.Fatalf("%d = %v", i, results[i])
		}
	}
}
