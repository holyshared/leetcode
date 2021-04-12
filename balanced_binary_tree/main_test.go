package main

import "testing"

// [3,9,20,null,null,15,7]
func TestA(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	if isBalanced(root) != true {
		t.Fatal("oops!!")
	}

}

// [1,2,2,3,null,null,3,4,null,null,4]
/*
             1
       2         2
     3  null   null  3
   4    null     null  4

*/
func TestB(t *testing.T) {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}

	if isBalanced(root) != false {
		t.Fatal("oops!!")
	}

}
