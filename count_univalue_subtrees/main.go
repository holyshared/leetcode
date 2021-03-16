package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func _uni(root *TreeNode, curr int) (bool, int) {
	if root.Left == nil && root.Right == nil {
		curr++
		return true, curr
	}
	unival := true
	ok := false

	if root.Left != nil {
		ok, curr = _uni(root.Left, curr) 
		unival = ok && unival && root.Left.Val == root.Val
	}

	if root.Right != nil {
		ok, curr = _uni(root.Right, curr) 
		unival = ok && unival && root.Right.Val == root.Val
	}
	if !unival {
		return false, curr
	}
	curr++
	return true, curr
}

func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, count := _uni(root, 0)
	return count
}
