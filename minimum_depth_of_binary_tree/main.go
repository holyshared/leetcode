package main

import (
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minOf(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func _minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	min := math.MaxInt
	if root.Left != nil {
		min = minOf(_minDepth(root.Left), min)
	}

	if root.Right != nil {
		min = minOf(_minDepth(root.Right), min)
	}
	return min + 1
}

func minDepth(root *TreeNode) int {
	return _minDepth(root)
}