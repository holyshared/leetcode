package main

import "testing"

/*
        1
   2        7

3    4         8
  5    6     9


        F
   B        G

A    D        I
  C    E     H

325461798
*/

func TestA(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}}, Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}}}}

	actual := inorderTraversal(root)
	expected := []int{3, 2, 5, 4, 6, 1, 7, 9, 8}

	for i, v := range actual {
		if v != expected[i] {
			t.Fatalf("%d: %d != %d", i, v, expected[i])
		}
	}
}
