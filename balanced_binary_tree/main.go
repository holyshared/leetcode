package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(node *TreeNode) int {
	if node == nil {
		return -1
	}
	lh, rh := height(node.Left), height(node.Right)
	if lh > rh {
		return lh + 1
	} else {
		return rh + 1
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := abs(height(root.Left)-height(root.Right)) < 2

	return diff && isBalanced(root.Left) && isBalanced(root.Right)
}
