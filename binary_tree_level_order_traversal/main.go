package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _levelOrder(node *TreeNode, level int, levels [][]int) [][]int {
	if node == nil {
		return levels
	}
	curr := len(levels)
	if curr == level {
		levels = append(levels, []int{})
	}

	levels[level] = append(levels[level], node.Val)
	if node.Left != nil {
		levels = _levelOrder(node.Left, level+1, levels)
	}
	if node.Right != nil {
		levels = _levelOrder(node.Right, level+1, levels)
	}

	return levels
}

func levelOrder(root *TreeNode) [][]int {
	return _levelOrder(root, 0, [][]int{})
}
