package main

import "testing"

/*
     4
   /   \
  2     6
 / \   /
3   1 5
*/
func TestA(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}}}

	result := addOneRow(root, 1, 2)

	if result.Val != 4 {
		t.Fatalf("%d", result.Val)
	}
	if result.Left.Val != 1 {
		t.Fatalf("%d", result.Left.Val)
	}
	if result.Left.Left.Val != 2 {
		t.Fatalf("%d", result.Left.Left.Val)
	}
	if result.Left.Left.Left.Val != 3 {
		t.Fatalf("%d", result.Left.Left.Left.Val)
	}
	if result.Left.Left.Right.Val != 1 {
		t.Fatalf("%d", result.Left.Left.Right.Val)
	}

	if result.Right.Val != 1 {
		t.Fatalf("%d", result.Right.Val)
	}
	if result.Right.Right.Val != 6 {
		t.Fatalf("%d", result.Right.Right.Val)
	}
	if result.Right.Right.Left.Val != 5 {
		t.Fatalf("%d", result.Right.Right.Left.Val)
	}
}

/*
4
/
2
/ \
3   1
*/
func TestB(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}}

	result := addOneRow(root, 1, 3)

	if result.Val != 4 {
		t.Fatalf("%d", result.Val)
	}
	if result.Left.Val != 2 {
		t.Fatalf("%d", result.Left.Val)
	}
	if result.Left.Left.Val != 1 {
		t.Fatalf("%d", result.Left.Left.Val)
	}
	if result.Left.Left.Left.Val != 3 {
		t.Fatalf("%d", result.Left.Left.Left.Val)
	}
	if result.Left.Right.Val != 1 {
		t.Fatalf("%d", result.Left.Right.Val)
	}
	if result.Left.Right.Right.Val != 1 {
		t.Fatalf("%d", result.Left.Right.Right.Val)
	}

	/*
	   	4
	   	/
	   2
	    / \
	   1   1
	   /     \
	   3       1
	*/

}

func TestC(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}

	result := addOneRow(root, 5, 4)

	if result.Val != 1 {
		t.Fatalf("%d", result.Val)
	}
	if result.Left.Val != 2 {
		t.Fatalf("%d", result.Left.Val)
	}
	if result.Left.Left.Val != 4 {
		t.Fatalf("%d", result.Left.Left.Val)
	}
	if result.Left.Left.Left.Val != 5 {
		t.Fatalf("%d", result.Left.Left.Left.Val)
	}
	if result.Left.Left.Right.Val != 5 {
		t.Fatalf("%d", result.Left.Left.Right.Val)
	}
	if result.Right.Val != 3 {
		t.Fatalf("%d", result.Right.Val)
	}

	/*
			       1
			   	 / \
			   	2   3
			   /
		    4
			 / \
		  5   5
	*/

}
