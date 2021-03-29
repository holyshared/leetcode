package main

import "testing"

func TestA(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	actual := flipMatchVoyage(root, []int{2, 1})

	if actual[0] != -1 {
		t.Fatal("oops!!")
	}

}

func TestB(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	actual := flipMatchVoyage(root, []int{1, 3, 2})

	if actual[0] != 1 {
		t.Fatal("oops!!")
	}
}

func TestC(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	actual := flipMatchVoyage(root, []int{1, 2, 3})

	if len(actual) > 0 {
		t.Fatal("oops!!")
	}
}
