package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _inorderTraversal(node *TreeNode, results []int) []int {
	if node == nil {
		return results
	}
	if node.Left != nil {
		results = _inorderTraversal(node.Left, results)
	}
	results = append(results, node.Val)
	if node.Right != nil {
		results = _inorderTraversal(node.Right, results)
	}
	return results
}

func inorderTraversal(root *TreeNode) []int {
	return _inorderTraversal(root, []int{})
}
