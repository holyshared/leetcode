package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _postorderTraversal(node *TreeNode, results []int) []int {
	if node == nil {
		return results
	}
	if node.Left != nil {
		results = _postorderTraversal(node.Left, results)
	}
	if node.Right != nil {
		results = _postorderTraversal(node.Right, results)
	}
	return append(results, node.Val)
}

func postorderTraversal(root *TreeNode) []int {
	return _postorderTraversal(root, []int{})
}
