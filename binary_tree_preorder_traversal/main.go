package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _preorderTraversal(node *TreeNode, results []int) []int {
	if node == nil {
		return results
	}
	results = append(results, node.Val)

	if node.Left != nil {
		results = _preorderTraversal(node.Left, results)
	}
	if node.Right != nil {
		results = _preorderTraversal(node.Right, results)
	}
	return results
}

func preorderTraversal(root *TreeNode) []int {
	return _preorderTraversal(root, []int{})
}
