package main

import "testing"

func Test123_5(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	if hasPathSum(root, 5) != false {
		t.Fatal("oops!!")
	}
}

func Test123_3(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	if hasPathSum(root, 3) != true {
		t.Fatal("oops!!")
	}
}

func Test123_4(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	if hasPathSum(root, 4) != true {
		t.Fatal("oops!!")
	}
}

func TestNil_0(t *testing.T) {
	if hasPathSum(nil, 0) != false {
		t.Fatal("oops!!")
	}
}

func Test12(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}

	if hasPathSum(root, 1) != false {
		t.Fatal("oops!!")
	}
}

func Test1(t *testing.T) {
	root := &TreeNode{Val: 1}

	if hasPathSum(root, 1) != true {
		t.Fatal("oops!!")
	}
}
