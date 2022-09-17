package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _tree2str(node *TreeNode) string {
	if node == nil {
		return ""
	}

	r := strconv.Itoa(node.Val)

	lv := _tree2str(node.Left)
	rv := _tree2str(node.Right)

	if lv == "" && rv != "" {
		return r + "()(" + rv + ")"
	} else if lv != "" && rv == "" {
		return r + "(" + lv + ")"
	} else if lv == "" && rv == "" {
		return r
	}

	return r + "(" + lv + ")(" + rv + ")"
}

func tree2str(root *TreeNode) string {
	return _tree2str(root)
}
