package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _goodNodes(node *TreeNode, routeMax int) int {
	if node == nil {
		return 0
	}

	goodCount := 0
	nextMax := routeMax
	if node.Val >= routeMax {
		goodCount += 1
		nextMax = node.Val
	}

	if node.Left != nil {
		goodCount += _goodNodes(node.Left, nextMax)
	}
	if node.Right != nil {
		goodCount += _goodNodes(node.Right, nextMax)
	}
	return goodCount
}

func goodNodes(root *TreeNode) int {
	return _goodNodes(root, math.MinInt32)
}
