package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _right(root *TreeNode, level int, results []int) []int {
  if level == len(results) {
		results = append(results, root.Val)
	}

	if root.Right != nil {
		results = _right(root.Right, level + 1, results)
	} 
	if root.Left != nil {
		results = _right(root.Left, level + 1, results)
	} 

	return results
}

func rightSideView(root *TreeNode) []int {
	results := []int{}
	if root == nil {
		return results
	}
	return _right(root, 0, results)
}
